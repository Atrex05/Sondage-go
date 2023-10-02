package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func convertResponseToInt(responseNumber string) (int, error) {
	if responseNumber != "" {
		return strconv.Atoi(responseNumber)
	}
	return 0, nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(session.Values, "authenticated")
	delete(session.Values, "username")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteSondageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("trying to delete sondage")
	index := r.PostFormValue("index")
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		log.Fatal(err)
	}
	userID, err := getLoggedInUserID(w, r)
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("SELECT createur_id FROM sondages WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var verif_userID int
	stmt.QueryRow(indexInt).Scan(&verif_userID)
	if userID != verif_userID {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	deleteSondage(indexInt)
	http.Redirect(w, r, "/sondage_list", http.StatusSeeOther)
}

func execTemplate(w http.ResponseWriter, page string) {
	tmpl, err := template.ParseFiles("../static/" + page + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
