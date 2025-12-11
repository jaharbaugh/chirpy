package main

import(
	"net/http"
	"log"
	
)


func (cfg *apiConfig) handlerGetChirps (w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	allChirps, err := cfg.db.GetChirps(req.Context())
	if err != nil{
		log.Printf("Fetch Chirps error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "Error fetching chirps", err)
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range allChirps {
    	chirps = append(chirps, Chirp{
        ID:        dbChirp.ID,
        CreatedAt: dbChirp.CreatedAt,
        UpdatedAt: dbChirp.UpdatedAt,
        Body:      dbChirp.Body,
        UserID:    dbChirp.UserID,
    	})
	}

	respondJSON(w, http.StatusOK, chirps)
}	