package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var db *sql.DB

type ChoiceCount struct {
	ChoiceCount int `json:"choiceCount"`
}

var choiceCount ChoiceCount

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db = nil
	err := error(nil)
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connexion à la base de données réussie.")

	err = CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	go deleteExpiredSondages()
	load_programm()
}

func load_programm() error {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/ws", websocketHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/styles.css", staticHandler)
	http.HandleFunc("/signup", signupHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/create_sondage", createSondageHandler)
	http.HandleFunc("/sondage_list", sondageListHandler)
	http.HandleFunc("/sondage/", SondageHandler)
	http.HandleFunc("/sondage/styles.css", staticHandler)
	http.HandleFunc("/sondage_results/", sondageResultsHandler)
	http.HandleFunc("/reset_password", ResetPasswordHandler)
	http.HandleFunc("/password_recovery", PasswordRecoveryHandler)
	http.HandleFunc("/delete_sondage", deleteSondageHandler)
	http.HandleFunc("/update_choice_count", updateChoiceCountHandler)
	log.Println("Serveur en écoute sur le port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
