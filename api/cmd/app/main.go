package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmettoguz/Youtube_MP3_Converter/internal/health"
	"github.com/ahmettoguz/Youtube_MP3_Converter/internal/youtube"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})

	http.HandleFunc("/health", health.HealthHandler)

	// New endpoint: /download?url=YOUTUBE_URL
	http.HandleFunc("/download", youtube.DownloadHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
