package main

import (
	"buhtigexa.snippetbox.com/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (a *Application) createSnippetController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Snippet
	if err := json.Unmarshal(bytes, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a.createSnippet(data)
	w.WriteHeader(http.StatusCreated)
}

func (a *Application) getById(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.UserAgent())

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}
