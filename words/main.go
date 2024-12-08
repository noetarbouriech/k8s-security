package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

//go:embed nouns.json
var nouns []byte

//go:embed verbs.json
var verbs []byte

type Words struct {
	nouns []string
	verbs []string
}

func NewWords() (*Words, error) {
	var n []string
	var v []string
	if err := json.Unmarshal(nouns, &n); err != nil {
		return nil, fmt.Errorf("failed to decode nouns JSON: %w", err)
	}
	if err := json.Unmarshal(verbs, &v); err != nil {
		return nil, fmt.Errorf("failed to decode verbs JSON: %w", err)
	}
	return &Words{
		nouns: n,
		verbs: v,
	}, nil
}

func RandomWord(words []string) string {
	if len(words) == 0 {
		return "no words available"
	}
	return words[rand.Intn(len(words))]
}

func (words *Words) nounsHandler(w http.ResponseWriter, r *http.Request) {
	randomNoun := RandomWord(words.nouns)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(randomNoun))
}

func (words *Words) verbsHandler(w http.ResponseWriter, r *http.Request) {
	randomVerb := RandomWord(words.verbs)
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(randomVerb))
}

func main() {
	words, err := NewWords()
	if err != nil {
		fmt.Println("Error initializing words:", err)
		return
	}

	http.HandleFunc("/nouns", words.nounsHandler)
	http.HandleFunc("/verbs", words.verbsHandler)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
