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
	
	cfg :=  &apiConfig{}
	cfg.fileserverHits.Store(0)
	cfg.db = database.New(db)
	cfg.Platform = os.Getenv("PLATFORM")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	multiplexer.HandleFunc("GET /api/healthz", handlerHealthz)
	multiplexer.HandleFunc("GET /admin/metrics", cfg.handlerMetrics)
	multiplexer.HandleFunc("POST /admin/reset", cfg.handlerReset)

	multiplexer.HandleFunc("POST /api/chirps", cfg.handlerChirpsCreate)
	multiplexer.HandleFunc("GET /api/chirps", cfg.handlerGetChirps)
	multiplexer.HandleFunc("GET /api/chirps/{chirpID}", cfg.handlerGetChirpsByID)
	
	multiplexer.HandleFunc("POST /api/users", cfg.handlerUsers)
	multiplexer.HandleFunc("PUT /api/users", cfg.handlerUpdateUser)
	multiplexer.HandleFunc("POST /api/login", cfg.handlerLogin)
	multiplexer.HandleFunc("POST /api/refresh", cfg.handlerRefresh)
	multiplexer.HandleFunc("POST /api/revoke", cfg.handlerRevoke)

	var server http.Server
	server.Handler = multiplexer
	server.Addr = ":8080"

	dir := http.Dir(".")
	fileServer := http.FileServer(dir)
	strippedFileServer := http.StripPrefix("/app", fileServer)
	multiplexer.Handle("/app/", cfg.middlewareMetricsInc(strippedFileServer))

	log.Fatal(server.ListenAndServe())

}









