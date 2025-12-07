package main

import(
	"net/http"
	//"fmt"
	"sync/atomic"
	"log"
	//"encoding/json"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func main(){

	multiplexer := http.NewServeMux()
	
	cfg:=  &apiConfig{}
	cfg.fileserverHits.Store(0)

	multiplexer.HandleFunc("GET /api/healthz", handlerHealthz)
	multiplexer.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	multiplexer.HandleFunc("POST /admin/reset", cfg.handlerReset)
	multiplexer.HandleFunc("POST /api/validate_chirp", handlerValidate)
	
	var server http.Server
	server.Handler = multiplexer
	server.Addr = ":8080"

	dir := http.Dir(".")
	fileServer := http.FileServer(dir)
	strippedFileServer := http.StripPrefix("/app", fileServer)
	multiplexer.Handle("/app/", cfg.middlewareMetricsInc(strippedFileServer))

	log.Fatal(server.ListenAndServe())

}









