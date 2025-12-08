package main

import(
	"encoding/json"
	"net/http"
	
)


func decode (req *http.Request) (Chirp, error) {

	decoder := json.NewDecoder(req.Body)
	chirpBody := Chirp{}
	err := decoder.Decode(&chirpBody)
	if err != nil{
		return Chirp{}, err
	}

	return chirpBody, nil

}