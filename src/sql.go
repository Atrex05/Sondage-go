package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sondage"
	password = "bood7Ees"
	dbname   = "sondage_db"
)

func CreateVotesTable(db *sql.DB) error {
	var tableExists bool = true
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "votes").Scan(&tableExists)
	if err != nil {
		return err
	}
	if !tableExists {
		_, err := db.Exec("CREATE TABLE votes (id int, vote INT);")
		if err != nil {
			return err
		}
		fmt.Println("Table 'votes' créée avec succès.")
	} else {
		fmt.Println("La table 'votes' existe déjà.")
	}
	return nil
}

func CreateUsersTable(db *sql.DB) error {
	var tableExists bool = true
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)", "users").Scan(&tableExists)
	if err != nil {
		return err
	}
	if !tableExists {
		_, err := db.Exec("CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR(50) UNIQUE, password VARCHAR(100));")
		if err != nil {
			return err
		}
		fmt.Println("Table 'users' créée avec succès.")
	} else {
		fmt.Println("La table 'users' existe déjà.")
	}
	return nil
}

func CreateTable(db *sql.DB) error {
	err := CreateVotesTable(db)
	if err != nil {
		return err
	}
	err = CreateUsersTable(db)
	if err != nil {
		return err
	}
	return nil
}

func deleteSondage(sondageID int) {
	var tableName string
	err := db.QueryRow("SELECT titre FROM sondages WHERE id = $1", sondageID).Scan(&tableName)
	if err != nil {
		log.Fatal(err)
	}
	tableName = getTableName(tableName)
	stmt := fmt.Sprintf("DROP TABLE %s", tableName)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	stmt = "DELETE FROM sondages WHERE id = $1"
	_, err = db.Exec(stmt, sondageID)
	if err != nil {
		log.Fatal(err)
	}
}
