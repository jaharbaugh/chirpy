package main

import(
	"net/http"
	"log"
	"github.com/google/uuid"
	"sort"
)


func (cfg *apiConfig) handlerGetChirps (w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	idStr := req.URL.Query().Get("author_id")
		
	if idStr != "" {

		userID, err := uuid.Parse(idStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Could not parse author ID", err)
			return
		}

		userChirps, err := cfg.db.GetChirpsByUserID(req.Context(), userID)
		if err != nil{
			log.Printf("Could not fetch User's chirps error: %v", err)
			respondWithError(w, http.StatusInternalServerError, "Error fetching chirps", err)
			return	
		}

		chirps := []Chirp{}
		for _, dbChirp := range userChirps {
    		chirps = append(chirps, Chirp{
        	ID:        dbChirp.ID,
        	CreatedAt: dbChirp.CreatedAt,
        	UpdatedAt: dbChirp.UpdatedAt,
        	Body:      dbChirp.Body,
        	UserID:    dbChirp.UserID,
    		})
		}

			sortString := req.URL.Query().Get("sort")
			if sortString == "asc" || sortString == ""{
				sort.Slice(chirps, func(i, j int) bool {
					return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
				})
			}else if sortString == "desc"{
			sort.Slice(chirps, func(i, j int) bool {
				return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
			})
		}

		respondJSON(w, http.StatusOK, chirps)
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

	sortString := req.URL.Query().Get("sort")
	if sortString == "asc" || sortString == ""{
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
		})
	}else if sortString == "desc"{
		sort.Slice(chirps, func(i, j int) bool {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		})
	}

	respondJSON(w, http.StatusOK, chirps)
}	