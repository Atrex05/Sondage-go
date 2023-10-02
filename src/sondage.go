package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type voteData struct {
	Choix1      string
	Choix2      string
	Choix3      string
	Choix4      string
	ChoiceCount int
	Error       string
	Titre       string
	IsUser      bool
}

func getVoteData(w http.ResponseWriter, r *http.Request, id int) voteData {
	data := voteData{}
	stmt, err := db.Prepare(`
		SELECT choix1, choix2, choix3, choix4, nb_choices, titre
		FROM sondages 
		WHERE id = $1
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&data.Choix1,
		&data.Choix2,
		&data.Choix3,
		&data.Choix4,
		&data.ChoiceCount,
		&data.Titre,
	)
	if err != nil {
		log.Fatal(err)
	}
	if !userIsLoggedIn(w, r) {
		data.Error = "Vous ne pouvez pas voter, vous n'êtes pas connecté"
	} else if userHasVoted(w, r, id) {
		data.Error = "Vous avez déjà voté"
	}
	return data
}

func isOneInSondageIDs(id int) bool {
	stmt, err := db.Prepare("SELECT EXISTS (SELECT 1 FROM sondages WHERE id = $1)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var exists bool = false
	err = stmt.QueryRow(id).Scan(&exists)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return exists
}

func SondageHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")
	var id string

	if len(parts) == 3 && parts[1] == "sondage" && parts[2] != "styles.css" {
		id = parts[2]
	} else {
		http.Error(w, "Page non trouvée", http.StatusNotFound)
		return
	}
	if id == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}
	if !isOneInSondageIDs(idInt) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		if !userIsLoggedIn(w, r) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else if userHasVoted(w, r, idInt) {
			http.Redirect(w, r, "/sondage_results/"+id, http.StatusSeeOther)
			return
		} else {
			responseNumber := r.FormValue("choix")
			fmt.Println("responsenum ====" + responseNumber)
			intResponseNumber, err := convertResponseToInt(responseNumber)
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := db.Prepare("SELECT titre FROM sondages WHERE id = $1")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			var tableName string
			err = stmt.QueryRow(idInt).Scan(&tableName)
			if err != nil {
				log.Fatal(err)
			}
			tableName = getTableName(tableName)
			userID, err := getLoggedInUserID(w, r)
			if err != nil {
				log.Fatal(err)
			}
			stmt, err = db.Prepare(fmt.Sprintf("INSERT INTO %s (user_id, vote_id) VALUES ($1, $2)", tableName))
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(userID, intResponseNumber)
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/sondage_results/"+id, http.StatusSeeOther)
		}
	}

	data := getVoteData(w, r, idInt)
	data.IsUser = false

	if userIsCreator(w, r, idInt) {
		data.IsUser = true
	}
	tmpl, err := template.ParseFiles("../static/sondage.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
