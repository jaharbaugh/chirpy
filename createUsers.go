package main

import(
	"net/http"
	"log"
	"github.com/jaharbaugh/chirpy/internal/auth"
	"github.com/jaharbaugh/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	newUser, err := decode[NewUserParams](req)
	if err != nil{
		log.Printf("CreateUser error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}

	hashedPW, err := auth.HashPassword(newUser.Password)
		if err != nil{
			log.Printf("CreateUser error: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Password Error", err)
		}
	var createUser database.CreateUserParams
		createUser.Email = newUser.Email
		createUser.HashedPassword = hashedPW

	res, err := cfg.db.CreateUser(req.Context(), createUser)
	if err != nil{
		log.Printf("CreateUser error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}
	
	var user User
	user.ID = res.ID
	user.CreatedAt = res.CreatedAt
	user.UpdatedAt = res.UpdatedAt
	user.Email = res.Email

	respondJSON(w, http.StatusCreated, user)


}