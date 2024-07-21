package server

import (
	"log"
	"net/http"
	"os"
	"path"
	"sharemail/internal/config"
	/* "sharemail/internal/server/components" */)

func Start() {
	// Load environment variables from .env file
	staticFilesPath := path.Join(config.AppConfig["ROOT_PATH"], "/web/static/")
	config.LoadEnv()
	port := os.Getenv("APP_PORT")
	/* component := components.Hello("World") */

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/ping", ping)
	mux.HandleFunc("/api/v1/email/link", getEmailLink)
	mux.Handle(
		"/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir(staticFilesPath))),
	)
	/* mux.Handle("/tests/component", templ.Handler(component)) */
	handler := addMiddlware(mux, logMiddleware)

	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe("localhost:"+port, handler)
	if err != nil {
		panic(err)
	}
}
