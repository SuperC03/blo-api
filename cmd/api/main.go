package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/superc03/blo-api/config"
	_ "github.com/superc03/blo-api/docs"
	"go.uber.org/zap"
)

//	@title		Banana Lounge API
//	@version	1.0

func main() {
	cfg, err := config.ParseEnv()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gotify := config.NewGotifyClient(cfg)

	log, err := config.NewLogger(cfg, gotify)
	if err != nil {
		fmt.Println(err.Error())
	}

	r := chi.NewRouter()

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port),
		Handler: r,
	}
	serverCtx, serverCtxStop := context.WithCancel(context.Background())
	serverStopSig := make(chan os.Signal, 1)
	signal.Notify(serverStopSig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-serverStopSig
		shutdownCtx, shutdownCtxStop := context.WithTimeout(serverCtx, 30*time.Second)
		defer shutdownCtxStop()
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful exit failed, forcing exit")
			}
		}()
		if err = server.Shutdown(shutdownCtx); err != nil {
			log.Fatal("unable to shutdown server", zap.Error(err))
		}
		serverCtxStop()
	}()

	log.Info("starting server", zap.String("addr", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("unable to start server", zap.Error(err))
	}
	<-serverCtx.Done()
}
