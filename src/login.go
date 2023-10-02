package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store = sessions.NewCookieStore([]byte("HJEETDFb5Uoq1zb8LibT57B/oRSSV/SJkxgjRMEWoL04I9wXYUyad8w7n5bHlUIk"))

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func checkCredentials(w http.ResponseWriter, r *http.Request, db *sql.DB, username string, password string) bool {
	id := -1
	stmt, err := db.Prepare("SELECT id FROM users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.QueryRow(username).Scan(&id)
	if id == -1 {
		return false
	}

	stmt, err = db.Prepare("SELECT password FROM users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var hashedPassword string
	err = stmt.QueryRow(username).Scan(&hashedPassword)
	if err != nil {
		log.Fatal(err)
	}

	isValid := checkPasswordHash(password, hashedPassword)
	if isValid {
		fmt.Println("Login successful")
		return true
	}
	fmt.Println("Invalid password")
	return false
}

func createSession(w http.ResponseWriter, r *http.Request, username string) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authenticated"] = true
	session.Values["username"] = username

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if checkCredentials(w, r, db, username, password) {
			createSession(w, r, username)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			time.Sleep(1 * time.Second)
			tmpl, err := template.ParseFiles("../static/login.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			errMessage := "Nom d'utilisateur ou mot de passe incorrect"
			tmpl.Execute(w, errMessage)
			return
		}
	}
	time.Sleep(1 * time.Second)
	execTemplate(w, "login")
}
