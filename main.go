package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
)

func main() {
	l := log.New(os.Stdout, "gh-ingestor", log.Ldate|log.Ltime|log.Lshortfile)

	hook, err := github.New(github.Options.Secret(os.Getenv("WEBHOOK_SECRET")))
	if err != nil {
		l.Fatal(err)
	}

	eventHandler := &EventHandler{hook: hook, log: l}

	r := gin.Default()
	r.POST("/", eventHandler.HandleEvents)

	server := &http.Server{
		Addr:     ":3000",
		Handler:  r,
		ErrorLog: l,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			l.Fatalf("Server error: %v", err)
		}
	}()

	gracefulShutdown(server)
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}
}
