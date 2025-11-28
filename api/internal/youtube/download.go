package youtube

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/kkdai/youtube/v2"
)

// DownloadAndConvertMP3 downloads YouTube audio and embeds thumbnail as album art
func DownloadAndConvertMP3(url, outputDir string) (string, error) {
	client := youtube.Client{}

	video, err := client.GetVideo(url)
	if err != nil {
		return "", err
	}

	// Sanitize title for filename
	title := video.Title
	mp3File := filepath.Join(outputDir, sanitizeFileName(title)+".mp3")
	tmpThumbnail := filepath.Join(outputDir, "thumbnail_temp.jpg")

	// Download thumbnail if available
	if len(video.Thumbnails) > 0 {
		thumbURL := video.Thumbnails[len(video.Thumbnails)-1].URL // highest quality
		if err := downloadThumbnail(thumbURL, tmpThumbnail); err != nil {
			fmt.Println("Warning: failed to download thumbnail:", err)
			tmpThumbnail = "" // ignore if failed
		}
	}

	// Get audio format
	audioFormats := video.Formats.Type("audio/mp4")
	if len(audioFormats) == 0 {
		return "", fmt.Errorf("no audio format found")
	}
	format := audioFormats[0]
	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		return "", err
	}

	tmpAudio := filepath.Join(outputDir, "audio_temp.m4a")

	// Save audio stream to temporary file
	audioFile, err := os.Create(tmpAudio)
	if err != nil {
		return "", err
	}
	defer audioFile.Close()
	if _, err := io.Copy(audioFile, stream); err != nil {
		return "", err
	}

	// Embed thumbnail using ffmpeg
	if tmpThumbnail != "" {
		cmd := exec.Command(
			"ffmpeg",
			"-i", tmpAudio,
			"-i", tmpThumbnail,
			"-map", "0:a",
			"-map", "1:v",
			"-c:a", "libmp3lame",
			"-b:a", "192k",
			"-id3v2_version", "3",
			"-metadata:s:v", "title=\"Album cover\"",
			"-metadata:s:v", "comment=\"Cover (front)\"",
			"-y", mp3File,
		)
		if err := cmd.Run(); err != nil {
			fmt.Println("Warning: failed to embed thumbnail:", err)
			os.Rename(tmpAudio, mp3File)
		} else {
			os.Remove(tmpAudio)
		}
		os.Remove(tmpThumbnail)
	} else {
		// Just rename temp audio
		os.Rename(tmpAudio, mp3File)
	}

	fmt.Println("Downloaded MP3 with cover:", mp3File)
	return mp3File, nil
}

// downloadThumbnail downloads the thumbnail from URL
func downloadThumbnail(url, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

// sanitizeFileName removes invalid characters
func sanitizeFileName(name string) string {
	re := regexp.MustCompile(`[<>:"/\\|?*]`)
	return re.ReplaceAllString(name, "_")
}
