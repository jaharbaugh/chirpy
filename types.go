package main

import (
	"sync/atomic"
	"github.com/jaharbaugh/chirpy/internal/database"
	"time"
	"github.com/google/uuid"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db *database.Queries
	Platform string
	
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

type NewUser struct {
	Email string `json:"email"`
}

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	 `json:"updated_at"`
	Email string `json:"email"`
}