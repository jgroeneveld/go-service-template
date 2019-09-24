package app

import (
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"net/http"
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
