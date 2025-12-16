package main

import(
	"net/http"
//	"context"
)

func (cfg *apiConfig) handlerReset (w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	if cfg.Platform != "dev" {
    	respondWithError(w, http.StatusForbidden, "forbidden", nil)
    	return
	}
	
	err := cfg.db.ResetUsers(req.Context())
	if err!= nil{
		respondWithError(w, http.StatusInternalServerError, "Reset Failed", err)
		return 
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)

}