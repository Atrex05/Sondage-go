package main

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	testDBHost     = "localhost"
	testDBPort     = 5432
	testDBUser     = "sondage"
	testDBPassword = "bood7Ees"
	testDBName     = "sondage_db"
)

func TestCreateVotestable(t *testing.T) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		testDBHost, testDBPort, testDBUser, testDBPassword, testDBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connexion à la base de données réussie.")

	err = CreateVotesTable(db)
	if err != nil {
		fmt.Println("ERROR: Erreur lors de la création de la table 'votes':", err)
	} else {
		fmt.Println("SUCCESS: Table 'votes' créée avec succès.")
	}
}
func TestCreateUsersTable(t *testing.T) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		testDBHost, testDBPort, testDBUser, testDBPassword, testDBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = CreateUsersTable(db)
	if err != nil {
		fmt.Println("ERROR: Erreur lors de la création de la table 'users':", err)
	} else {
		fmt.Println("SUCCESS: Table 'users' créée avec succès.")
	}
}

func TestCreateTable(t *testing.T) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		testDBHost, testDBPort, testDBUser, testDBPassword, testDBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = CreateUsersTable(db)
	if err != nil {
		fmt.Println("ERROR: Erreur lors de la création des tables:", err)
	} else {
		fmt.Println("SUCCESS: Table créées avec succès.")
	}
}
