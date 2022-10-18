package jokebuilder

import (
	"task/internal/jokefetcher"
	"task/internal/namefetcher"
)

type JokeFetcher interface {
	FetchJoke(firstName, lastName, category string) (*jokefetcher.JokeResponse, error)
}

type NameFetcher interface {
	FetchName() (*namefetcher.NameResponse, error)
}

type JokeBuilder struct {
	jokeFetcher JokeFetcher
	nameFetcher NameFetcher
}

func NewJokeBuilder(jokeFetcher JokeFetcher, nameFetcher NameFetcher) *JokeBuilder {
	return &JokeBuilder{
		jokeFetcher: jokeFetcher,
		nameFetcher: nameFetcher,
	}
}

func (j *JokeBuilder) BuildJoke(first, last, category string) (string, error) {
	
}
