package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Эндпоинт для метрик Prometheus
	http.Handle("/metrics", promhttp.Handler())

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
