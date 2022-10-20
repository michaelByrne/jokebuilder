package main

import (
	"log"
	"net/http"
	"task/app"
	"task/internal/jokebuilder"
	"task/internal/jokefetcher"
	"task/internal/namefetcher"
	"task/internal/transport"
)

func main() {
	client := &http.Client{}
	jokeFetcher := jokefetcher.NewJokeFetcher(client)
	nameFetcher := namefetcher.NewNameFetcher(client)
	jokeBuilder := jokebuilder.NewJokeBuilder(jokeFetcher, nameFetcher)

	handler := transport.NewHandler(jokeBuilder)

	jokeApp := app.NewApp(handler)

	log.Fatal(jokeApp.Run())
}
