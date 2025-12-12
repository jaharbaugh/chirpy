package main

import(
	"net/http"
	"log"
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

	var verifiedUser User
	verifiedUser.ID = dbUser.ID
	verifiedUser.CreatedAt = dbUser.CreatedAt
	verifiedUser.UpdatedAt = dbUser.UpdatedAt
	verifiedUser.Email = dbUser.Email

	respondJSON(w, http.StatusOK, verifiedUser)

}