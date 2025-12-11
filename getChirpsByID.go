package main
import (
	"net/http"
	"log"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerGetChirpsByID (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	idStr := req.PathValue("chirpID")

	id, err := uuid.Parse(idStr)
	if err != nil {
    	respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
    	return
	}

	dbChirp, err := cfg.db.GetChirpByID(req.Context(), id)
	if err != nil{
		log.Printf("Fetch Chirps error: %v", err)
		respondWithError(w, http.StatusNotFound, "Error fetching chirps", err)
		return
	}

	var chirp Chirp
	chirp.ID = dbChirp.ID
	chirp.CreatedAt = dbChirp.CreatedAt
	chirp.UpdatedAt = dbChirp.UpdatedAt
	chirp.Body = dbChirp.Body
	chirp.UserID = dbChirp.UserID 

	respondJSON(w, http.StatusOK, chirp)

}