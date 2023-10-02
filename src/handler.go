package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	ButtonText  string
	IsConnected bool
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	loggedUsername, ok := session.Values["username"].(string)
	if !ok {
		loggedUsername = ""
	}
	data := PageData{
		IsConnected: false,
	}
	if loggedUsername == "" {
		data.ButtonText = "No user logged in"
		data.IsConnected = false
	} else {
		data.ButtonText = loggedUsername
		data.IsConnected = true
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/styles.css" {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "../static/styles.css")
		return
	}
	if r.URL.Path == "sondage/styles.css" {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "../static/styles.css")
		return
	}
	if r.URL.Path == "/sondage_results/styles.css" {
		w.Header().Set("Content-Type", "text/css")
		http.ServeFile(w, r, "../static/sondage_results/styles.css")
		return
	}
}

func updateChoiceCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&choiceCount); err != nil {
		http.Error(w, "Erreur lors de la lecture du JSON", http.StatusBadRequest)
		return
	}

	createurID, err := getLoggedInUserID(w, r)
	if err != nil {
		log.Fatal(err)
	}
	query := `UPDATE sondages SET nb_choices = $1 WHERE id = $2`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(choiceCount.ChoiceCount, createurID)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
