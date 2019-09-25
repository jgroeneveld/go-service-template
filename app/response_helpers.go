package app

import (
	"encoding/json"
	"net/http"
)

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

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}