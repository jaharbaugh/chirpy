package main

import(
	"net/http"
	"fmt"
	"sync/atomic"
	"log"
)

type apiConfig struct {
	fileServerHits atomic.Int32
}

func main(){

	multiplexer := http.NewServeMux()
	
	cfg:=  &apiConfig{}
	cfg.fileServerHits.Store(0)

	multiplexer.HandleFunc("GET /healthz", handlerHealthz)
	multiplexer.HandleFunc("GET /metrics", cfg.handlerMetrics )
	multiplexer.HandleFunc("POST /reset", cfg.handlerReset)
	
	var server http.Server
	server.Handler = multiplexer
	server.Addr = ":8080"

	dir := http.Dir(".")
	fileServer := http.FileServer(dir)
	strippedFileServer := http.StripPrefix("/app", fileServer)
	multiplexer.Handle("/app/", cfg.middlewareMetricsInc(strippedFileServer))

	log.Fatal(server.ListenAndServe())

}

func handlerHealthz (w http.ResponseWriter, req *http.Request){ 
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) handlerReset (w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	cfg.fileServerHits.Store(0)

}

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %v\n", cfg.fileServerHits.Load())))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}