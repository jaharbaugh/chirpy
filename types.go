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
	JWTSecret string
	
}

type NewChirpParams struct {
    Body string `json:"body"`
	//UserID uuid.UUID `json:"user_id"`
}

type Chirp struct{
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	 `json:"updated_at"`
	Body string `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

type ResponseInvalid struct {
	Error string `json:"error"`	
}

type ResponseValid struct{
	Valid bool   `json:"valid"`
	Cleaned string `json:"cleaned_body"`
}

type NewUserParams struct {
	Email string `json:"email"`
	Password string `json:"password"`
	ExpiresInSeconds time.Duration `json:"expires_in_secpnds"`
}

type User struct {
	ID uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time	 `json:"updated_at"`
	Email string `json:"email"`
	HashedPassword string `json:"-"`
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`

}

type NewAccessToken struct{
	Token string `json:"token"`
}