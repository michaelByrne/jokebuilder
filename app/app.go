package app

import (
	"log"
	"net/http"
	"task/internal/jokebuilder"
	"task/internal/transport"
)

type App struct {
	jokeBuilder *jokebuilder.JokeBuilder
	handler     *transport.Handler
}

func NewApp(jokeBuilder *jokebuilder.JokeBuilder, handler *transport.Handler) *App {
	return &App{
		jokeBuilder: jokeBuilder,
		handler:     handler,
	}
}

func (a *App) Run() error {
	http.HandleFunc("/", a.handler.Handle)

	log.Println("Starting server on port 8080")

	return http.ListenAndServe(":8080", nil)
}
