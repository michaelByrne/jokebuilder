package transport

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakeJokeBuilder struct {
	called bool
}

func (f *fakeJokeBuilder) BuildJoke(category string) (string, error) {
	f.called = true
	return "knock knock", nil
}

func TestHandler_Handle(t *testing.T) {
	jokeBuilder := &fakeJokeBuilder{}
	handler := NewHandler(jokeBuilder)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	h := http.HandlerFunc(handler.Handle)

	h.ServeHTTP(recorder, req)

	require.Equal(t, jokeBuilder.called, true)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "knock knock", recorder.Body.String())
}
