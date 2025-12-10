package main

import(
	"encoding/json"
	"net/http"
	
)


func decode[T any](req *http.Request) (T, error) {
	var v T
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&v)
	return v, err
}