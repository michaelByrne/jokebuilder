package transport

import (
	"net/http"
	"task/internal/jokebuilder"
)

type Handler struct {
	jokeBuilder jokebuilder.JokeBuilder
}

func NewHandler(jokeBuilder jokebuilder.JokeBuilder) *Handler {
	return &Handler{
		jokeBuilder: jokeBuilder,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, _ *http.Request) {
	joke, err := h.jokeBuilder.BuildJoke("nerdy")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(joke))
}
