package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Sondage struct {
	Name            string
	ID              int
	Description     string
	Expiration_date string
	Creation_date   string
}

func sondageListHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/sondage_list.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("SELECT id, titre, description, date_creation, date_expiration FROM sondages")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sondages []Sondage
	for rows.Next() {
		var id int
		var name string
		var description string
		var dateCreation time.Time
		var dateExpiration time.Time

		err := rows.Scan(&id, &name, &description, &dateCreation, &dateExpiration)
		if err != nil {
			log.Fatal(err)
		}
		dateCreationString := "Créé le " + dateCreation.Format("02/01 à 15h04")
		dateExpirationString := "Expire le " + dateExpiration.Format("02/01 à 15h04")
		sondage := Sondage{
			ID:              id,
			Name:            name,
			Description:     description,
			Creation_date:   dateCreationString,
			Expiration_date: dateExpirationString,
		}

		sondages = append(sondages, sondage)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, sondages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
