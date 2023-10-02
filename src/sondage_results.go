package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type sondageData struct {
	Choix1 int
	Choix2 int
	Choix3 int
	Choix4 int
}

type voteChoice struct {
	Vote1       string
	Vote2       string
	Vote3       string
	Vote4       string
	ChoiceCount int
}

func getSondageData(w http.ResponseWriter, r *http.Request, id int) sondageData {
	data := sondageData{}
	var tableName string
	db.QueryRow("SELECT titre FROM sondages WHERE id = $1", id).Scan(&tableName)
	tableName = getTableName(tableName)

	data.Choix1 = getCountForVoteID(tableName, 1)
	data.Choix2 = getCountForVoteID(tableName, 2)
	data.Choix3 = getCountForVoteID(tableName, 3)
	data.Choix4 = getCountForVoteID(tableName, 4)

	return data
}

func getCountForVoteID(tableName string, voteID int) int {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM " + tableName + " WHERE vote_id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var count int
	if err := stmt.QueryRow(voteID).Scan(&count); err != nil {
		log.Fatal(err)
	}

	return count
}

func getVoteChoice(w http.ResponseWriter, r *http.Request, id int) voteChoice {
	data := voteChoice{}
	stmt, err := db.Prepare(`
        SELECT choix1, choix2, choix3, choix4, nb_choices 
        FROM sondages 
        WHERE id = $1
    `)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(
		&data.Vote1,
		&data.Vote2,
		&data.Vote3,
		&data.Vote4,
		&data.ChoiceCount,
	)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func sondageResultsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/sondage_results.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	path := r.URL.Path
	parts := strings.Split(path, "/")
	var id string

	if len(parts) == 3 && parts[2] != "styles.css" {
		id = parts[2]
	}
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	sondageID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	data := getVoteChoice(w, r, sondageID)
	fmt.Println("DATA: ", data.ChoiceCount)
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
