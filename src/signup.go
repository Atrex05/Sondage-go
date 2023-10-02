package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username     string
	Password     string
	Mail_address string
	IsConnected  bool
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func createUser(w http.ResponseWriter, db *sql.DB, username string, password string, Mail string) (bool, error) {
	id := -1
	stmt, err := db.Prepare("SELECT id FROM users WHERE username = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	stmt.QueryRow(username).Scan(&id)
	if id != -1 {
		tmpl, err := template.ParseFiles("../static/signup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false, err
		}
		errMessage := "Nom d'utilisateur déjà utilisé"
		tmpl.Execute(w, errMessage)
		return false, nil
	}
	db.QueryRow("SELECT id FROM users WHERE mail_address = $1", Mail).Scan(&id)
	if id != -1 {
		tmpl, err := template.ParseFiles("../static/signup.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false, err
		}
		errMessage := "Adresse mail déjà utilisée"
		tmpl.Execute(w, errMessage)
		return false, nil
	}
	stmt, err = db.Prepare("INSERT INTO users (username, password, mail_address) VALUES ($1, $2, $3)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, password, Mail)
	if err != nil {
		log.Fatal(err)
	}
	return true, err
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		mail_address := r.FormValue("mail")
		if !isValidEmail(mail_address) {
			tmpl, err := template.ParseFiles("../static/signup.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			errMessage := "Adresse mail invalide"
			tmpl.Execute(w, errMessage)
			return
		}
		password, err := hashPassword(password)
		if err != nil {
			log.Fatal(err)
		}
		user := User{
			Username:     username,
			Password:     password,
			Mail_address: mail_address,
		}
		isValid, err := createUser(w, db, user.Username, user.Password, user.Mail_address)
		if err != nil {
			log.Fatal(err)
		}
		if isValid {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		return
	}
	execTemplate(w, "signup")
}
