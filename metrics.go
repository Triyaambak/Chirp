package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) HandlerMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`
		<html>
			<body>
				<h1>Welcome, Chirpy Admin</h1>
				<p>Chirpy has been visited %d times!</p>
			</body>
		</html>
	`, cfg.fileServeHits)))
}

func (cfg *apiConfig) ResetMetrics(w http.ResponseWriter, r *http.Request) {
	cfg.fileServeHits = 0
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server hits set to 0"))
}

func (cfg *apiConfig) MiddlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileServeHits++
		next.ServeHTTP(w, r)
	})
}
