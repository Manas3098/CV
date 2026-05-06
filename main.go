package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	mux := http.NewServeMux()

	// Serve static files
	mux.Handle("/static/", http.FileServer(http.FS(staticFiles)))

	// Main route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticFiles, "static/index.html")
	})

	log.Println("🔐 Portfolio server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
