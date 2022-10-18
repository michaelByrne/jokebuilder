package jokebuilder

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"task/internal/jokefetcher"
	"task/internal/namefetcher"
	"testing"
)

type fakeJokeFetcher struct {
	called             bool
	calledWithFirst    string
	calledWithLast     string
	calledWithCategory string
	response           *jokefetcher.JokeResponse
	err                error
}

func (f *fakeJokeFetcher) FetchJoke(firstName, lastName, category string) (*jokefetcher.JokeResponse, error) {
	f.called = true
	f.calledWithFirst = firstName
	f.calledWithLast = lastName
	f.calledWithCategory = category
	return f.response, f.err
}

type fakeNameFetcher struct {
	called   bool
	response *namefetcher.NameResponse
	err      error
}

func (f *fakeNameFetcher) FetchName() (*namefetcher.NameResponse, error) {
	f.called = true
	return f.response, f.err
}

func TestJokeBuilder_BuildJoke(t *testing.T) {
	t.Run("should return error if name fetcher returns error", func(t *testing.T) {
		nameFetcher := &fakeNameFetcher{
			err:      errors.New("some error"),
			response: nil,
		}
		jokeFetcher := &fakeJokeFetcher{
			response: nil,
			err:      nil,
		}
		jokeBuilder := NewJokeBuilder(jokeFetcher, nameFetcher)

		_, err := jokeBuilder.BuildJoke("dev")
		require.Error(t, err)

		assert.True(t, nameFetcher.called)
		assert.False(t, jokeFetcher.called)
	})

	t.Run("should return error if joke fetcher returns error", func(t *testing.T) {
		nameFetcher := &fakeNameFetcher{
			err: nil,
			response: &namefetcher.NameResponse{
				First: "John",
				Last:  "Doe",
			},
		}
		jokeFetcher := &fakeJokeFetcher{
			response: nil,
			err:      errors.New("some error"),
		}
		jokeBuilder := NewJokeBuilder(jokeFetcher, nameFetcher)

		_, err := jokeBuilder.BuildJoke("dev")
		require.Error(t, err)

		assert.True(t, nameFetcher.called)
		assert.True(t, jokeFetcher.called)
		assert.Equal(t, "John", jokeFetcher.calledWithFirst)
		assert.Equal(t, "Doe", jokeFetcher.calledWithLast)
		assert.Equal(t, "dev", jokeFetcher.calledWithCategory)
	})

	t.Run("should successfully return a joke", func(t *testing.T) {
		nameFetcher := &fakeNameFetcher{
			err: nil,
			response: &namefetcher.NameResponse{
				First: "John",
				Last:  "Doe",
			},
		}
		jokeFetcher := &fakeJokeFetcher{
			response: &jokefetcher.JokeResponse{
				Type: "joke",
				Value: struct {
					Categories []string `json:"categories"`
					ID         int      `json:"id"`
					Joke       string   `json:"joke"`
				}{
					Categories: nil,
					ID:         0,
					Joke:       "some joke",
				},
			},
			err: nil,
		}
		jokeBuilder := NewJokeBuilder(jokeFetcher, nameFetcher)

		joke, err := jokeBuilder.BuildJoke("dev")
		require.NoError(t, err)

		assert.True(t, nameFetcher.called)
		assert.True(t, jokeFetcher.called)
		assert.Equal(t, "John", jokeFetcher.calledWithFirst)
		assert.Equal(t, "Doe", jokeFetcher.calledWithLast)
		assert.Equal(t, "dev", jokeFetcher.calledWithCategory)
		assert.Equal(t, "some joke", joke)
	})
}
