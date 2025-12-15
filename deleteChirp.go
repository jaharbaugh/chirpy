package main

import(
	"net/http"
	"log"
	"github.com/jaharbaugh/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodDelete {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	token, err := auth.GetBearerToken(req.Header)
	if err != nil {
    	respondWithError(w, http.StatusUnauthorized, "Missing or invalid auth header", err)
    	return
	}

	tokenUserID, err := auth.ValidateJWT(token, cfg.JWTSecret)
	if err != nil {
    	respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
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
		log.Printf("Fetch Chirp error: %v", err)
		respondWithError(w, http.StatusNotFound, "Error fetching chirps", err)
		return
	}

	if dbChirp.UserID != tokenUserID {
		respondWithError(w, http.StatusForbidden, "Not authorized to delete this chirp", nil)
		return
	}

	err = cfg.db.DeleteChirp(req.Context(), dbChirp.ID)
	if err != nil{
		log.Printf("Chirp not deleted : %v", err)
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}	

