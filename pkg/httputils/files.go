package httputils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var (
	extension = map[string]string{
		"image/png":       ".png",
		"image/jpeg":      ".jpeg",
		"image/webp":      ".webp",
		"application/pdf": ".pdf",
	}

	ImagesMime = []string{"image/png", "image/jpeg", "image/webp"}
	PDFmime    = "application/pdf"
)

// SaveFile - Saves file in the provided path. the key indicates the name of
// the file that is indicated in the sent form. the acceptedMIMETypes determine
// if the file type. SaveFile returns file path along any errors
func SaveFile(r *http.Request, key, path string, acceptedMIMETypes ...string) (string, error) {
	path = filepath.Clean(path)

	if len(acceptedMIMETypes) == 0 {
		return "", errors.New("provide at least one MIME type")
	}

	file, _, err := r.FormFile(key)
	if err != nil {
		return "", err
	}

	defer file.Close()

	contentType, err := detectType(file)
	if err != nil {
		return "", err
	}

	var found bool
	for _, t := range acceptedMIMETypes {
		if t == contentType {
			found = true
			break
		}
	}

	if !found {
		return "", errors.New("file not supported")
	}

	ext, ok := extension[contentType]
	if !ok {
		return "", errors.New("extension doesn't exist")
	}

	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	fileName, err := generateRandomString(10)
	if err != nil {
		return "", err
	}

	fullPath := filepath.Join(path, fmt.Sprintf("%v%v", fileName, ext))

	dst, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(fullPath)
		return "", err
	}

	return fullPath, nil
}

// detectType - detects the MIME Type for a file. if it can't determine it, it
// returns application/octet-stream
func detectType(file multipart.File) (string, error) {
	buffer := make([]byte, 512)
	if _, err := file.Read(buffer); err != nil {
		return "", err
	}

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	return http.DetectContentType(buffer), nil
}

func generateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes), nil
}
