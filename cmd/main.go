package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"

	"github-ingestor-go/internal"
	"github-ingestor-go/pkg/memphis"
)

func main() {
	l := log.New(os.Stdout, "gh-ingestor", log.Ldate|log.Ltime|log.Lshortfile)

	hook, err := github.New(github.Options.Secret(os.Getenv("WEBHOOK_SECRET")))
	if err != nil {
		l.Fatal(err)
	}

	pc, err := getProducerConfig()
	if err != nil {
		l.Fatal("Unable to construct ProducerConfig, env variables are probably missing: ", err.Error())
	}
	eventHandler := internal.NewEventHandler(hook, l, pc)

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

func getProducerConfig() (*memphis.ProducerConfig, error) {
	memphis_acc_id, err := strconv.Atoi(os.Getenv("MEMPHIS_ACCOUNT_ID"))
	if err != nil {
		return &memphis.ProducerConfig{}, err
	}

	return memphis.NewProducerConfig(memphis_acc_id, os.Getenv("MEMPHIS_HOST"), os.Getenv("MEMPHIS_USERNAME"), os.Getenv("MEMPHIS_PASSWORD")), nil
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
