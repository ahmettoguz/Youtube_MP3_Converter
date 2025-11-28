package youtube

import (
    "fmt"
    "os/exec"
    "regexp"

    "github.com/kkdai/youtube/v2"
)

// DownloadAndConvertMP3 downloads a YouTube video and converts it directly to MP3
func DownloadAndConvertMP3(url, outputDir string) (string, error) {
    client := youtube.Client{}

    video, err := client.GetVideo(url)
    if err != nil {
        return "", err
    }

    // sanitize title for filename
    title := video.Title
    mp3File := fmt.Sprintf("%s/%s.mp3", outputDir, sanitizeFileName(title))

    // get best audio format
    format := video.Formats.WithAudioChannels()[0]

    // use ffmpeg with stdin from youtube stream
    stream, _, err := client.GetStream(video, &format)
    if err != nil {
        return "", err
    }

    cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-vn", "-ab", "128k", "-ar", "44100", "-y", mp3File)
    cmd.Stdin = stream

    if err := cmd.Run(); err != nil {
        return "", err
    }

    fmt.Println("Downloaded MP3:", mp3File)
    return mp3File, nil
}

func sanitizeFileName(name string) string {
    re := regexp.MustCompile(`[<>:"/\\|?*]`)
    return re.ReplaceAllString(name, "_")
}
