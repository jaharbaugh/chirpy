package main

import(
	"net/http"
	"github.com/jaharbaugh/chirpy/internal/database"
	"github.com/google/uuid"
	"log"
	"github.com/jaharbaugh/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeChirpyRed (w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	keyString, err := auth.GetAPIKey(req.Header)
	if err != nil{
		log.Printf("Could not retrieve API Key")
		respondWithError(w, http.StatusUnauthorized,"Could not retrieve API Key", err)
		return
	}
	if keyString != cfg.PolkaKey{
		log.Printf(keyString, cfg.PolkaKey)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	webhook, err := decode[Webhook](req)
	if err != nil{
		log.Printf("Failed to decode webhook : %v", err)
		respondWithError(w, http.StatusBadRequest, "Could not decode webhook", err)
		return
	}

	if webhook.Event != "user.upgraded"{
		w.WriteHeader(http.StatusNoContent)
		return
	}

	id, err := uuid.Parse(webhook.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user_id in webhook", err)
		return
	}	

	var params database.SetUserChirpyRedParams
	params.ID = id
	params.IsChirpyRed = true

	err = cfg.db.SetUserChirpyRed(req.Context(), params)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	w.WriteHeader(http.StatusNoContent)
}