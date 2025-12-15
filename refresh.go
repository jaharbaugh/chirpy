package main

import(
	"net/http"
	"github.com/jaharbaugh/chirpy/internal/auth"
	"time"
	"log"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	refreshToken, err := auth.GetBearerToken(req.Header)
	if err != nil{
		log.Printf("Refresh error: %v", err)
		respondWithError(w, http.StatusBadRequest, "Token not found", err)
		return
	}

	dbUser, err := cfg.db.GetUserByToken(req.Context(), refreshToken)
	if err != nil{
		log.Printf("Refresh error: %v", err)
		respondWithError(w, http.StatusUnauthorized, "User not found", err)
		return
	}

	accessToken, err := auth.MakeJWT(dbUser.ID, cfg.JWTSecret, time.Hour)
	if err != nil{
		log.Printf("Error creating new access token: %v", err)
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
    	return
	}

	/*var newToken NewAccessToken
	newToken.Token = accessToken

	respondJSON(w, http.StatusOK, newToken.Token)
	*/
	type response struct {
        Token string `json:"token"`
    }

    respondJSON(w, http.StatusOK, response{
        Token: accessToken,
    })
}
