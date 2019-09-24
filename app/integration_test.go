package app

import (
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"github.com/jgroeneveld/trial/th"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootResponse(t *testing.T) {
	server := routes()

	response := Get(server, "/")

	assert.MustBeEqual(t, http.StatusOK, response.Code)
	AssertJSONSchema(t, response.Body, schema.Map{
		"Foo": "Hello World",
	})
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

func AssertJSONSchema(t *testing.T, r io.Reader, matcher schema.Matcher) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
	}
}

func MustMatchJSONSchema(t *testing.T, r io.Reader, matcher schema.Matcher) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, err.Error())
		t.FailNow()
	}
}
