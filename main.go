package main

import(
	"net/http"
	//"fmt"
	//"sync/atomic"
	"log"
	//"encoding/json"
	//"strings"
	_ "github.com/lib/pq"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/jaharbaugh/chirpy/internal/database"
)

func main(){
	godotenv.Load()
	dbURL:= os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil{
		panic(err)
	}
	
	multiplexer := http.NewServeMux()
	
	cfg:=  &apiConfig{}
	cfg.fileserverHits.Store(0)
	cfg.db = database.New(db)

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









