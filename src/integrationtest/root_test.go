package integrationtest

import (
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	router := NewTestSetup().Router()

	response := Get(router, "/")

	assert.MustBeEqual(t, http.StatusOK, response.Code)
	assert.JSONSchema(t, response.Body, schema.Map{
		"Foo": "Hello World",
	})
}
