package main

import (
	"fmt"
	"net/http"

	"github.com/ahmettoguz/Youtube_MP3_Converter/internal/health"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go Running...")
	})

	// Health endpoint returns JSON
	http.HandleFunc("/health", health.HealthHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
