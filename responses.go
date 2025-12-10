package main

import (
	"net/http"
	"encoding/json"
	"log"
)

func responseValid() (ResponseValid, int){
	res := ResponseValid{
		Valid: true,
		Cleaned: "",
	}
 
	return res, http.StatusOK
}

func responseInvalid(err string) (ResponseInvalid, int){
	res := ResponseInvalid{
		Error: err,
	}
	 
	return res, http.StatusBadRequest 
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
    w.WriteHeader(status)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, status int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if status > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}

	respondJSON(w, status, ResponseInvalid{
		Error: msg,
	})
}