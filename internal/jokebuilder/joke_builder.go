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

func (j *JokeBuilder) BuildJoke(category string) (string, error) {
	name, err := j.nameFetcher.FetchName()
	if err != nil {
		return "", err
	}

	joke, err := j.jokeFetcher.FetchJoke(name.First, name.Last, category)
	if err != nil {
		return "", err
	}

	return joke.Value.Joke, nil
}
