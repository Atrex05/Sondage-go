package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type formData struct {
	Titre          string
	Description    string
	Choix1         string
	Choix2         string
	Choix3         string
	Choix4         string
	CreateurID     int
	CurrentTime    time.Time
	DateExpiration time.Time
}

func tableExists(db *sql.DB, tableName string) (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
		AND table_name = $1
	)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}
	var exists bool
	err = stmt.QueryRow(tableName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func getTableName(name string) string {
	name = strings.Replace(name, " ", "_", -1)
	if len(name) > 63 {
		name = name[:63]
	}
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	name = re.ReplaceAllString(name, "")
	return name
}

func createResultTable(db *sql.DB, tableName string) error {
	tableName = getTableName(tableName)
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS ` + tableName + ` (
		user_id INT,
		vote_id INT
	)`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func getFormData(w http.ResponseWriter, r *http.Request) formData {
	data := formData{
		Titre:       r.FormValue("titre"),
		Description: r.FormValue("description"),
		Choix1:      r.FormValue("choix1"),
		Choix2:      r.FormValue("choix2"),
		Choix3:      r.FormValue("choix3"),
		Choix4:      r.FormValue("choix4"),
		CreateurID:  1,
	}
	createurID, err := getLoggedInUserID(w, r)
	data.CreateurID = createurID
	if err != nil {
		log.Fatal(err)
	}
	dureeHours, err := strconv.Atoi(r.FormValue("duree-hours"))
	if err != nil {
		log.Fatal(err)
	}
	dureeMinutes, err := strconv.Atoi(r.FormValue("duree-minutes"))
	if err != nil {
		log.Fatal(err)
	}
	if dureeHours == 0 && dureeMinutes == 0 {
		dureeHours = 4
	}
	currentTime := time.Now()
	duree := time.Duration(dureeHours)*time.Hour + time.Duration(dureeMinutes)*time.Minute
	dateExpiration := currentTime.Add(duree)
	data.CurrentTime = currentTime
	data.DateExpiration = dateExpiration
	return data
}

func createSondage(d formData, choiceCount ChoiceCount, w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO sondages (titre, description, choix1, choix2, choix3, choix4, nb_choices, createur_id, date_creation, date_expiration) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.Titre, d.Description, d.Choix1, d.Choix2, d.Choix3, d.Choix4, choiceCount.ChoiceCount, d.CreateurID, d.CurrentTime, d.DateExpiration)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createSondageHandler(w http.ResponseWriter, r *http.Request) {
	if !userIsLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if r.Method == http.MethodPost {
		d := getFormData(w, r)
		isTable, err := tableExists(db, d.Titre)
		if err != nil {
			log.Fatal(err)
		}
		if isTable {
			errMsg := "Un sondage avec ce titre existe déjà"
			tmpl, err := template.ParseFiles("../static/create_sondage.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, errMsg)
			return
		}
		createResultTable(db, d.Titre)
		createSondage(d, choiceCount, w, r)
		return
	}

	execTemplate(w, "create_sondage")
}
