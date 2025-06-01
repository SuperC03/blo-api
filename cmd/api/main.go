package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superc03/blo-api/api/routers"
	"github.com/superc03/blo-api/config"
	"github.com/superc03/blo-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

//	@title			Banana Lounge API
//	@version		1.0
//	@description	Server for accessing "public" data and internal reports

// @contact.name API Support
// @contact.email banana@colclark.net

// @basepath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY

func main() {
	serverCtx, serverCtxStop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer serverCtxStop()

	cfg, err := config.ParseEnv()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	docs.SwaggerInfo.Host = cfg.PublicURL
	docs.SwaggerInfo.BasePath = "/"

	gotify := config.NewGotifyClient(cfg)

	log, err := config.NewLogger(cfg, gotify)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db, err := config.NewDBConnection(serverCtx, cfg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	e := echo.New()
	e.Pre(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "docs")
		},
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.GET("/docs/*", echoSwagger.WrapHandler)
	routers.NewHealthRouter(e.Group("/health"), db, log, cfg, gotify, nil)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Hostname, cfg.Port),
		Handler: e,
	}

	go func() {
		log.Info("server starting 🎉", zap.String("addr", server.Addr))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Warn("error shutting down server", zap.Error(err))
		}
	}()
	<-serverCtx.Done()
	shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCtxCancel()
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Warn("error shutting down server", zap.Error(err))
	}
	log.Info("server shutdown 👋")

	<-serverCtx.Done()
}
