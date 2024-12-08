package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func fetchWord(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func sentenceHandler(w http.ResponseWriter, r *http.Request) {
	// Get the URLs from environment variables
	nounURL := os.Getenv("NOUN_URL")
	verbURL := os.Getenv("VERB_URL")

	if nounURL == "" || verbURL == "" {
		http.Error(w, "NOUN_URL or VERB_URL environment variable not set", http.StatusInternalServerError)
		return
	}

	// Fetch a noun and a verb
	noun, err := fetchWord(nounURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching noun: %v", err), http.StatusInternalServerError)
		return
	}

	verb, err := fetchWord(verbURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching verb: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the sentence
	sentence := fmt.Sprintf("The %s %s.", noun, verb)
	w.Write([]byte(sentence))
}

func main() {
	http.HandleFunc("/", sentenceHandler)

	log.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
