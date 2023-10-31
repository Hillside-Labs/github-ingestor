package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/webhooks/v6/github"
)

func main() {
	l := log.New(os.Stdout, "gh-ingestor", log.Ldate|log.Ltime|log.Lshortfile)
	hook, err := github.New(github.Options.Secret(""))
	if err != nil {
		l.Fatal(err)
	}

	eventHandler := &EventHandler{hook: hook, log: l}

	http.Handle("/events", eventHandler)
	http.ListenAndServe(":3000", nil)

}
