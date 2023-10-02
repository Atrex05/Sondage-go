package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var websocketConnections = make(map[int][]*websocket.Conn)

func cleanupWebsocketConnections(sondageID int) {
	updatedConnections := make([]*websocket.Conn, 0, len(websocketConnections[sondageID]))
	for _, conn := range websocketConnections[sondageID] {
		if _, _, err := conn.ReadMessage(); err != nil {
			continue
		}
		updatedConnections = append(updatedConnections, conn)
	}
	websocketConnections[sondageID] = updatedConnections
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Erreur lors de la mise à niveau de la connexion WebSocket:", err)
		return
	}
	defer conn.Close()
	sondageID := r.URL.Query().Get("sondageID")
	sondageIDInt, err := strconv.Atoi(sondageID)
	if err != nil {
		log.Fatal(err)
	}
	websocketConnections[sondageIDInt] = append(websocketConnections[sondageIDInt], conn)
	sendVoteUpdatesToClients(w, r)
	cleanupWebsocketConnections(sondageIDInt)
}

func sendVoteUpdatesToClients(w http.ResponseWriter, r *http.Request) {
	voteData := struct {
		Choix1 int
		Choix2 int
		Choix3 int
		Choix4 int
	}{}
	sondageID := r.URL.Query().Get("sondageID")
	sondageIDInt, err := strconv.Atoi(sondageID)
	if err != nil {
		log.Println("Erreur lors de la conversion de l'identifiant du sondage:", err)
		return
	}
	fmt.Println("Sondage ID:", sondageIDInt)
	sondageData := getSondageData(w, r, sondageIDInt)
	fmt.Println(sondageData)
	voteData.Choix1 = sondageData.Choix1
	voteData.Choix2 = sondageData.Choix2
	voteData.Choix3 = sondageData.Choix3
	voteData.Choix4 = sondageData.Choix4
	for _, conn := range websocketConnections[sondageIDInt] {
		err = conn.WriteJSON(voteData)
		if err != nil {
			log.Println("Erreur lors de l'envoi de données WebSocket:", err)
		}
	}
}
