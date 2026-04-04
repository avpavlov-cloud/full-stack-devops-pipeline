package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Full-Stack DevOps Pipeline (Golang edition)</h1>")
		fmt.Fprintf(w, "<p>Сервис запущен успешно!</p>")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Сервер запущен на порту %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
