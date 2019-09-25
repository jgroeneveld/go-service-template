package integrationtest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

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