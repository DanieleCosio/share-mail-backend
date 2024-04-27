package server

import (
	"encoding/json"
	"log"
	"net/http"
	"sharemail/internal/db"
)

func getEmailLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST.String() {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var messageData GetEmailLinkBody
	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if messageData.RequestAccountOwner == "" || messageData.MessageHtml == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("RequestAccountOwner: %s\n", messageData.RequestAccountOwner)
	log.Printf("MessageHtml: %s\n", messageData.MessageHtml)

	sqlConn, err := db.GetSqlConnection()
	if err != nil {
		log.Fatal(err)
	}

	id := 0
	query := `
		INSERT INTO email_link (owner_email, message_html) 
		VALUES ($1, $2)
		RETURNING id
	`
	err = sqlConn.QueryRow(
		query,
		messageData.RequestAccountOwner,
		messageData.MessageHtml,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}

	defer sqlConn.Close()
}

func ping(w http.ResponseWriter, r *http.Request) {
	jsonResponse := PingResponse{Success: true}
	jsonResponse.Dispatch(&w, jsonResponse)
}
