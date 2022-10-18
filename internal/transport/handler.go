package transport

import (
	"net/http"
)

type JokeBuilder interface {
	BuildJoke(category string) (string, error)
}

type Handler struct {
	jokeBuilder JokeBuilder
}

func NewHandler(jokeBuilder JokeBuilder) *Handler {
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
