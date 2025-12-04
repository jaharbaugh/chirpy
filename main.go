package main

import(
	"net/http"
)



func main(){

	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/healthz", handler)

	var server http.Server
	server.Handler = multiplexer
	server.Addr = ":8080"

	dir := http.Dir(".")
	fileServer := http.FileServer(dir)
	strippedFileServer := http.StripPrefix("/app", fileServer)
	multiplexer.Handle("/app/", strippedFileServer)

	server.ListenAndServe()

}

func handler (w http.ResponseWriter, req *http.Request){ 
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}