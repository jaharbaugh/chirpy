package main

import(
	"net/http"
	"log"
	"time"
	"github.com/jaharbaugh/chirpy/internal/auth"
	"github.com/jaharbaugh/chirpy/internal/database"
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

	token, err := auth.MakeJWT(dbUser.ID, cfg.JWTSecret, time.Hour)
	if err != nil{
		log.Printf("Could not make JWT Token: %v", err)
		respondWithError(w, http.StatusInternalServerError, "User not found", err)
		return
	}

	refreshTokenString, err := auth.MakeRefreshToken()
	if err != nil{	
		log.Printf("Could not make Refresh Token String: %v", err)
		respondWithError(w, http.StatusInternalServerError, "User not found", err)
		return
	}

	var newRefreshToken database.CreateRefreshTokenParams
	newRefreshToken.Token = refreshTokenString
	newRefreshToken.UserID = dbUser.ID

	refreshToken, err := cfg.db.CreateRefreshToken(req.Context(), newRefreshToken)
	if err != nil{
		log.Printf("Could not make Refresh Token: %v", err)
		respondWithError(w, http.StatusInternalServerError, "User not found", err)
		return
	}

	var verifiedUser User
	verifiedUser.ID = dbUser.ID
	verifiedUser.CreatedAt = dbUser.CreatedAt
	verifiedUser.UpdatedAt = dbUser.UpdatedAt
	verifiedUser.Email = dbUser.Email
	verifiedUser.Token = token
	verifiedUser.RefreshToken = refreshToken.Token
	verifiedUser.IsChirpyRed = dbUser.IsChirpyRed

	respondJSON(w, http.StatusOK, verifiedUser)

}