package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {
	const maxChirpLen int = 140
	type messageBody struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	body := messageBody{}
	err := decoder.Decode(&body)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode body", err)
		return
	}
	if len(body.Body) > maxChirpLen {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// Assuming the body passes length validation, check for profanity.
	body.Body = replaceBadWords(body.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: body.Body,
	})
}

func replaceBadWords(suspect string) string {
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	var goodString string

	// Case insensitive check
	words := strings.Split(suspect, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if _, ok := badWords[lowerWord]; ok {
			words[i] = "****"
		}
	}
	goodString = strings.Join(words, " ")
	return goodString
}
