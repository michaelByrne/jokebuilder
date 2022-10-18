package jokefetcher

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

const fakeJSONResponse = `{"type":"success","value":{"categories":["nerdy"],"id":1666127277,"joke":"\"It works on my machine\" always holds true for John Doe."}}`

type fakeHTTPClient struct {
	called     bool
	calledWith *http.Request
	response   *http.Response
	err        error
}

func (c *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.called = true
	c.calledWith = req
	return c.response, c.err
}

func TestJokeFetcher_FetchJoke(t *testing.T) {
	t.Run("it makes a GET request to the joke API", func(t *testing.T) {
		fakeHTTPResponse := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(fakeJSONResponse))),
		}

		client := &fakeHTTPClient{
			response: fakeHTTPResponse,
		}
		fetcher := NewJokeFetcher(client)

		_, err := fetcher.FetchJoke("John", "Doe", "nerdy")
		require.NoError(t, err)

		assert.True(t, client.called)

		assert.Equal(t, "GET", client.calledWith.Method)

		assert.Equal(t, "http://joke.loc8u.com:8888/joke?firstName=John&lastName=Doe&limitTo=nerdy", client.calledWith.URL.String())
	})

	t.Run("it returns an error if the request fails", func(t *testing.T) {
		client := &fakeHTTPClient{
			err: errors.New("some error"),
		}
		fetcher := NewJokeFetcher(client)

		_, err := fetcher.FetchJoke("John", "Doe", "nerdy")
		require.Error(t, err)

		assert.True(t, client.called)
	})

	t.Run("it returns a successful response", func(t *testing.T) {
		fakeHTTPResponse := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(fakeJSONResponse))),
		}

		client := &fakeHTTPClient{
			response: fakeHTTPResponse,
		}
		fetcher := NewJokeFetcher(client)

		joke, err := fetcher.FetchJoke("John", "Doe", "nerdy")
		require.NoError(t, err)

		assert.Equal(t, "\"It works on my machine\" always holds true for John Doe.", joke.Value.Joke)
	})
}
