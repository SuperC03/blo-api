package main

import (
	"context"
	"time"

	"github.com/superc03/blo-api/config"
	_ "github.com/superc03/blo-api/docs"
)

//	@title		Banana Lounge API
//	@version	1.0

func main() {
	gotify := config.NewGotifyClient(
		true,
		"https://gotify.colclark.net",
		"Awpa-aRO7gQWL8l",
	)
	go gotify.Send(context.Background(), "Hiy", "ya")
	l, _ := config.NewLogger(true, gotify)
	l.Fatal("Heyo")
	time.Sleep(5 * time.Second)
}
