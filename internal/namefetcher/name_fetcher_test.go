package namefetcher

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

const fakeJSONResponse = `{"first_name":"Hasina","last_name":"Tanweer"}`

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

func TestNameFetcher_FetchName(t *testing.T) {
	t.Run("it makes a GET request to the name API", func(t *testing.T) {
		fakeHTTPResponse := &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBuffer([]byte(fakeJSONResponse))),
		}

		client := &fakeHTTPClient{
			response: fakeHTTPResponse,
		}
		fetcher := NewNameFetcher(client)

		_, err := fetcher.FetchName()
		require.NoError(t, err)

		assert.True(t, client.called)

		assert.Equal(t, "GET", client.calledWith.Method)

		assert.Equal(t, "https://names.mcquay.me/api/v0/", client.calledWith.URL.String())
	})

	t.Run("it returns an error if the request fails", func(t *testing.T) {
		client := &fakeHTTPClient{
			err: errors.New("some error"),
		}
		fetcher := NewNameFetcher(client)

		_, err := fetcher.FetchName()
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
		fetcher := NewNameFetcher(client)

		name, err := fetcher.FetchName()
		require.NoError(t, err)

		assert.Equal(t, "Hasina", name.First)
		assert.Equal(t, "Tanweer", name.Last)
	})
}
