package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"sharemail/internal/db"

	"github.com/joho/godotenv"
)

type MessageData struct {
	RequestAccountOwner string `json:"requestAccountOwner"`
	MessageHtml         string `json:"messageHtml"`
}

func getEmailLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != POST.String() {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var messageData MessageData
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
}

func loadEnv() {
	filePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath = path.Join(filePath, ".env")

	err = godotenv.Load(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func Start() {
	// Load environment variables from .env file
	loadEnv()

	port := os.Getenv("APP_PORT")

	db.CreateConfig(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	// Routes
	http.HandleFunc("/api/v1/email/link", getEmailLink)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("web/"))))

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
