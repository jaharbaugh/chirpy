package main

import(
	"net/http"
	"log"
	"github.com/jaharbaugh/chirpy/internal/auth"
	"github.com/jaharbaugh/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateUser (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	bearerToken, err := auth.GetBearerToken(req.Header)
	if err != nil{
		log.Printf("Missing or malformed token error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Token error", err)
		return
	}

	tokenUserID, err := auth.ValidateJWT(bearerToken, cfg.JWTSecret)
	if err != nil {
		log.Printf("User not found with token error: %v", err)
    	respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
    	return
	}

	/*dbUser, err := cfg.db.GetUserByToken(req.Context(), tokenUserID)
	if err != nil{
		log.Printf("User Not Found Error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}*/

	userUpdates, err := decode[NewUserParams](req)
	if err != nil{
		log.Printf("Could not parse request: %v", err)
		respondWithError(w, http.StatusBadRequest, "Could not parse request", err)
		return
	}

	hash, err := auth.HashPassword(userUpdates.Password)
	if err != nil{
		log.Printf("Failed to hash password Error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Password not hashed", err)
		return
	}

	var updates database.UpdateUserParams
	updates.ID = tokenUserID
	updates.Email = userUpdates.Email
	updates.HashedPassword = hash

	err = cfg.db.UpdateUser(req.Context(), updates)
	if err != nil{
		log.Printf("Failed to update user error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to update user", err)
		return
	}

	updatedUser, err := cfg.db.GetUserByID(req.Context(), tokenUserID)
	if err != nil{
		log.Printf("Could not find user %v", err)
		respondWithError(w, http.StatusInternalServerError, "User not found", err)
	}

	response := User{
    ID:        updatedUser.ID,
    CreatedAt: updatedUser.CreatedAt,
    UpdatedAt: updatedUser.UpdatedAt,
    Email:     updatedUser.Email,
	}
	
	respondJSON(w, http.StatusOK, response)

}