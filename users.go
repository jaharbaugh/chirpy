package main

import(
	"net/http"
	//"encoding/json"
	"log"
)

func (cfg *apiConfig) handlerUsers(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	newUser, err := decode[NewUser](req)
	if err != nil{
		log.Printf("CreateUser error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}

	res, err := cfg.db.CreateUser(req.Context(), newUser.Email)
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