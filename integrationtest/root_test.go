package integrationtest

import (
	"github.com/jgroeneveld/go-cloud-run-template/app"
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	server := app.Routes()

	response := Get(server, "/")
	assert.MustBeEqual(t, http.StatusOK, response.Code)
	assert.JSONSchema(t, response.Body, schema.Map{
		"Foo": "Hello World",
	})
}
