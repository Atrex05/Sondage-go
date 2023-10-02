package main

import (
	"log"
	"time"
)

func deleteExpiredSondages() {
	for {
		currentTime := time.Now()
		query := `SELECT id FROM sondages WHERE date_expiration < $1`
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		rows, err := stmt.Query(currentTime)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var sondageID int
			if err := rows.Scan(&sondageID); err != nil {
				log.Fatal(err)
			}
			deleteSondage(sondageID)
		}
		time.Sleep(time.Hour)
		rows.Close()
	}
}
