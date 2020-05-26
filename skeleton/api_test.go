// Handler unit tests.
package skeleton

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerHandler(t *testing.T) {
	srv := NewServer(``)

	t.Run("GET /answer", func(t *testing.T) {
		got := struct {
			Result string
		}{}

		request := httptest.NewRequest(http.MethodGet, "/answer", nil)
		response := httptest.NewRecorder()
		srv.answer()(response, request)

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")
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

		request := httptest.NewRequest(http.MethodPost, "/answer", sent)
		response := httptest.NewRecorder()
		srv.answer()(response, request)

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, response.Result().StatusCode, "Status should be 200/OK")
		assert.Equal(t, "Hello, Stefano!", got.Result, "Result should be 'Hello, Stefano!'")
	})

}

func AsReader(t *testing.T, s interface{}) io.Reader {
	structBytes, err := json.Marshal(s)
	if err != nil {
		t.Fatal(err)
	}
	return bytes.NewReader(structBytes)
}
