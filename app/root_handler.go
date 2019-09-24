package app

import (
	"encoding/json"
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

func WriteJson(w http.ResponseWriter, v interface{}) error {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}

	return nil
}
