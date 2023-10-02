package main

import (
	"fmt"
	"log"
	"net/http"
)

func userIsLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	username, ok := session.Values["username"].(string)
	if username == "" || !ok {
		return false
	}
	return true
}

func getLoggedInUserID(w http.ResponseWriter, r *http.Request) (int, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return -1, err
	}
	username, ok := session.Values["username"].(string)
	if !ok {
		return -1, err
	}
	id := -1
	stmt, err := db.Prepare("SELECT id FROM users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(username).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id, nil
}

func userHasVoted(w http.ResponseWriter, r *http.Request, sondageID int) bool {
	userID, err := getLoggedInUserID(w, r)
	if err != nil {
		return false
	}
	var tableName string
	stmt, err := db.Prepare("SELECT titre FROM sondages WHERE id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(sondageID).Scan(&tableName)
	if err != nil {
		log.Fatal(err)
	}
	tableName = getTableName(tableName)
	var count int
	stmt, err = db.Prepare(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE user_id = $1", tableName))
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(userID).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		return false
	}
	return true
}

func userIsCreator(w http.ResponseWriter, r *http.Request, sondageID int) bool {
	userID, err := getLoggedInUserID(w, r)
	if err != nil {
		return false
	}
	query := `SELECT createur_id FROM sondages WHERE id = $1`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var creatorID int
	err = stmt.QueryRow(sondageID).Scan(&creatorID)
	if err != nil {
		log.Fatal(err)
	}
	if userID != creatorID {
		return false
	}
	return true
}
