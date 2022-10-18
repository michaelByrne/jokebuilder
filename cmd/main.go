package main

import (
	"log"
	"net/http"
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
	http.HandleFunc("/", handler.Handle)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
