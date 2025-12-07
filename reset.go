package main

import(
	"net/http"
)

func (cfg *apiConfig) handlerReset (w http.ResponseWriter, req *http.Request){
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
	
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	cfg.fileserverHits.Store(0)

}