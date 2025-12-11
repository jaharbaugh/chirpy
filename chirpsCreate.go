package main

import(
	"log"
	"net/http"
	"github.com/jaharbaugh/chirpy/internal/database"
)


func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	newChirp, err:= decode[NewChirpParams](req)
	if err != nil{
		log.Printf("CreateChirp error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}

	cleanedChirpBody, err := Validate(newChirp.Body)
	if err != nil {
        respondWithError(w, http.StatusBadRequest, "Error validating chirp", err)
        return
    }

	params := database.CreateChirpParams{
		Body: cleanedChirpBody,
		UserID: newChirp.UserID,
	}

	dbChirp, err := cfg.db.CreateChirp(req.Context(), params)
	if err != nil{
		log.Printf("CreateChirp error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Could not create chirp", err)
		return
	}
	
	var chirp Chirp
	chirp.ID = dbChirp.ID
	chirp.CreatedAt = dbChirp.CreatedAt
	chirp.UpdatedAt = dbChirp.UpdatedAt
	chirp.Body = dbChirp.Body
	chirp.UserID = dbChirp.UserID 

	respondJSON(w, http.StatusCreated, chirp)


}