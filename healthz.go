package main

import (
	"net/http"
)

func handlerHealthz (w http.ResponseWriter, req *http.Request){ 
	if req.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }
	
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}