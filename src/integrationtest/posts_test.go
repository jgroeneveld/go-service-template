package integrationtest

import (
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/assert"
	"net/http"
	"testing"
)

func TestPosts(t *testing.T) {
	router := NewTestSetup().Router()

	response := Get(router, "/posts/12")

	assert.MustBeEqual(t, http.StatusOK, response.Code)
	assert.JSONSchema(t, response.Body, schema.Map{
		"user_id": 42,
		"id":      12,
		"title":   "fake title",
		"body":    "fake body",
	})
}
