package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func ValidateJSON(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	chirp := Chirp{}
	err := json.NewDecoder(r.Body).Decode(&chirp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		data, _ := json.Marshal(map[string]string{"Error": "Something went wrong"})
		w.Write(data)
		return nil, err
	}

	if len(chirp.Body) > 140 {
		data, _ := json.Marshal(struct{ Error string }{Error: "Chirp is too long"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return nil, err
	}

	data := ReplaceWord(chirp.Body)
	res, _ := json.Marshal(map[string]string{"cleaned_body": data})
	return res, nil
}

func ReplaceWord(chirp string) (updatedChirp string) {
	var profaneWords = []string{"kerfuffle", "sharbert", "fornax"}
	for _, word := range profaneWords {
		chirp = strings.ReplaceAll(chirp, word, "****")
	}
	return chirp
}
