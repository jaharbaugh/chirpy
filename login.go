package main

import(
	"net/http"
	"log"
	"time"
	"github.com/jaharbaugh/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin (w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	unverifiedUser, err := decode[NewUserParams](req)
	if err != nil{
		log.Printf("User login error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}

	var expires time.Duration
	if unverifiedUser.ExpiresInSeconds > 0{
		expires = time.Duration(unverifiedUser.ExpiresInSeconds) * time.Second
		if expires > time.Hour{
			expires = time.Hour
		}
	}else{
		expires = time.Hour
	}

	dbUser, err := cfg.db.GetUserByEmail(req.Context(), unverifiedUser.Email)
	if err != nil{
		log.Printf("User login error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}

	match, err := auth.CheckPasswordHash(unverifiedUser.Password, dbUser.HashedPassword)
	if match != true{
		log.Printf("User login error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}

	token, err := auth.MakeJWT(dbUser.ID, cfg.JWTSecret, expires)
	if err != nil{
		log.Printf("Could not make JWT Token: %v", err)
		respondWithError(w, http.StatusInternalServerError, "User not found", err)
		return
	}

	var verifiedUser User
	verifiedUser.ID = dbUser.ID
	verifiedUser.CreatedAt = dbUser.CreatedAt
	verifiedUser.UpdatedAt = dbUser.UpdatedAt
	verifiedUser.Email = dbUser.Email
	verifiedUser.Token = token

	respondJSON(w, http.StatusOK, verifiedUser)

}