package app

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := routes()
	response := Get(server, "/")

	if status := response.Code; status != http.StatusOK {
		t.Errorf("actual: %v expected: %v", status, http.StatusOK)
	}
}

func Get(server http.Handler, url string) *httptest.ResponseRecorder {
	return SendRequest(server, "GET", url, nil)
}

func SendRequest(server http.Handler, method, url string, body io.Reader) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	return recorder
}
