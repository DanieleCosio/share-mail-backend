package server

import (
	"log"
	"net/http"
	"os"
	"path"
	"sharemail/internal/components"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

func loadEnv(projectRootPath string) {
	filePath := path.Join(projectRootPath, "/.env")
	err := godotenv.Load(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func Start() {
	// Load environment variables from .env file
	projectRootPath := getRootPath()
	staticFilesPath := path.Join(projectRootPath, "/web/static/")
	loadEnv(projectRootPath)
	port := os.Getenv("APP_PORT")
	component := components.Hello("World")

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/ping", ping)
	mux.HandleFunc("/api/v1/email/link", getEmailLink)
	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir(staticFilesPath))),
	)
	mux.Handle("/tests/component", templ.Handler(component))
	handler := addMiddlware(mux, logMiddleware)

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe("localhost:"+port, handler)
	if err != nil {
		log.Fatal(err)
	}
}
