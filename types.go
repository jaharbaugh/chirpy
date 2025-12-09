package main

import (
	"sync/atomic"
	"github.com/jaharbaugh/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
}

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