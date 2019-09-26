package app

import (
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	response := RootHandlerResponse{Foo: "Hello World"}

	err := WriteJson(w, response)
	if err != nil {
		panic(err)
	}
}

type RootHandlerResponse struct {
	Foo string
}
