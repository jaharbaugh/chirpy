package main

import (
	"net/http"
)

func responseValid() (ResponseValid, int){
	res := ResponseValid{
		Valid: true,
		Cleaned: "",
	}
 
	return res, http.StatusOK
}

func responseInvalid() (ResponseInvalid, int){
	res := ResponseInvalid{
		Error: "Chirp is too long",
	}
	 
	return res, http.StatusBadRequest 
}