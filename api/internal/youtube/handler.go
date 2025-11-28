package youtube

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// DownloadResponse represents JSON response
type DownloadResponse struct {
    Message string `json:"message"`
    MP3File string `json:"mp3_file,omitempty"`
    Error   string `json:"error,omitempty"`
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    if url == "" {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(DownloadResponse{
            Message: "error",
            Error:   "missing 'url' query parameter",
        })
        return
    }

    // download directly to MP3
    mp3File, err := DownloadAndConvertMP3(url, "output")
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        json.NewEncoder(w).Encode(DownloadResponse{
            Message: "error",
            Error:   err.Error(),
        })
        return
    }

    json.NewEncoder(w).Encode(DownloadResponse{
        Message: "success",
        MP3File: mp3File,
    })

    fmt.Println("Downloaded MP3:", mp3File)
}
