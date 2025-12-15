package main

import(
	"net/http"
	"log"
	"github.com/jaharbaugh/chirpy/internal/auth"
)

func (cfg apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil{
		log.Printf("Revoke error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Token not found", err)
		return
	}

	err = cfg.db.UpdateRefreshToken(req.Context(), refreshToken)
	if err != nil{
		log.Printf("Revoke error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Couldn't revoke session", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}