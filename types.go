package main

type Chirp struct {
    Body string `json:"body"`
}


type ResponseInvalid struct {
	Error string `json:"error"`	
}

type ResponseValid struct{
	Valid bool   `json:"valid"`
	Cleaned string `json:"cleaned_body"`
}