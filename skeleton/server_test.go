package skeleton

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswer(t *testing.T) {
	srv := NewServer(`/api`)

	t.Run("GET /answer", func(t *testing.T) {
		got := struct {
			Result string
		}{}

		request, err := http.NewRequest(http.MethodGet, "http://localhost:8000/api/answer", nil)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")

		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Hello!", got.Result, "Result should be 'Hello'")
	})

	t.Run("POST /answer", func(t *testing.T) {
		got := struct {
			Result string
		}{}
		sent := AsReader(t, struct {
			Name string
		}{
			Name: `Stefano`,
		})

		request, err := http.NewRequest(http.MethodPost, "http://localhost:8000/api/answer", sent)
		if err != nil {
			t.Fatal(err)
		}
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")

		err = json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, "Hello, Stefano!", got.Result, "Result should be 'Hello'")
	})
}
