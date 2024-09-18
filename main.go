package main

import (
	"log"
	"net/http"
	"sync"
)

type apiConfig struct {
	fileServeHits int
}

type Chirp struct {
	Body string `json:"body"`
}

type DB struct {
	path string
	mux  *sync.RWMutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

func main() {
	const PORT = "8080"

	apiCfg := &apiConfig{}
	dbCfg := &DB{path: "database.json"}

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: mux,
	}

	mux.Handle("GET /app", apiCfg.MiddlewareMetrics(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", HandlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)
	mux.HandleFunc("DELETE /api/reset", apiCfg.ResetMetrics)
	mux.HandleFunc("POST /api/chirps", dbCfg.CreateChirps)
	log.Printf("Serving on port: %s\n", PORT)
	log.Fatal(server.ListenAndServe())
}
