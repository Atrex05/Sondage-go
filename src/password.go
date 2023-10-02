package main

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net/http"

	"gopkg.in/gomail.v2"
)

var password_recovery string
var new_password_recovery string
var username_recovery string

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("mail")
		SendMailResetPassword(email, username, w, r)
	}

	tmpl, err := template.ParseFiles("../static/reset_password.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl.Execute(w, nil)
}

func PasswordRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../static/password_recovery.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if r.Method == http.MethodPost {
		password_recovery = r.FormValue("password")
		new_password_recovery = r.FormValue("new_password")
		if password_recovery == new_password_recovery {
			new_password_recovery, err = hashPassword(new_password_recovery)
			if err != nil {
				log.Fatal(err)
			}
			stmt, err := db.Prepare("UPDATE users SET password = $1 WHERE username = $2")
			if err != nil {
				log.Fatal(err)
			}
			defer stmt.Close()

			_, err = stmt.Exec(new_password_recovery, username_recovery)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			data := struct {
				Error string
			}{
				Error: "Les mots de passe ne correspondent pas.",
			}
			tmpl.Execute(w, data)
			return
		}
	}
	tmpl.Execute(w, nil)
}

func SendMail(from string, to string, subject string, body string, password string) error {
	smtpServer := "smtp-mail.outlook.com"
	smtpPort := 587
	message := gomail.NewMessage()

	message.SetHeader("From", from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/plain", body)

	dialer := gomail.NewDialer(smtpServer, smtpPort, from, password)

	if err := dialer.DialAndSend(message); err != nil {
		fmt.Println("Erreur lors de l'envoi du mail")
		panic(err)
	}

	println("E-mail envoyé avec succès.")
	return nil
}

func PasswordGenerator() string {
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+{}[]|:;<>,.?/~"
	passwordLength := 16 // Longueur du mot de passe souhaitée
	password := make([]byte, passwordLength)

	for i := 0; i < passwordLength; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			panic(err)
		}
		password[i] = characters[randomIndex.Int64()]
	}
	return string(password)
}
func SendMailResetPassword(email string, username string, w http.ResponseWriter, r *http.Request) {
	from := "comunesondage@outlook.com"
	to := email
	subject := "Réinitialisation de votre mot de passe"
	password := "J2t@#g+Mw9zK"
	token := PasswordGenerator()
	filepath := "/password_recovery" + token
	username_recovery = username

	body := "Bonjour " + username + ",\n\nVotre mot de passe a été réinitialisé.\n\nVotre nouveau mot de passe est disponible à l'adresse suivante: " + filepath + "\n\nCordialement,\nL'équipe Com-Une"
	http.HandleFunc(filepath, PasswordRecoveryHandler)
	err := SendMail(from, to, subject, body, password)
	if err != nil {
		log.Fatal(err)
	}
}
