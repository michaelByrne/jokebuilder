package app

import (
	"log"
	"net/http"
	"task/internal/transport"
)

type App struct {
	handler *transport.Handler
}

func NewApp(handler *transport.Handler) *App {
	return &App{
		handler: handler,
	}
}

func (a *App) Run() error {
	http.HandleFunc("/", a.handler.Handle)

	log.Println("Starting server on port 8080")

	return http.ListenAndServe(":8080", nil)
}
