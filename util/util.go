package util

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

func GetDefaultDownloadDirectory() string {
	usr, err := user.Current()
	if err != nil {
		return "./test"
	}
	t := filepath.Join(usr.HomeDir, "Downloads")
	return t
}

func SanitizeFileName(fileName string) string {
	sanitizedFileName := strings.TrimSpace(fileName)
	sanitizedFileName = strings.ReplaceAll(sanitizedFileName, " ", "_")

	removeSpecialCharsRegex := regexp.MustCompile(`[^\w\s._\&\,:-]`)
	sanitizedFileName = removeSpecialCharsRegex.ReplaceAllString(sanitizedFileName, "")

	return sanitizedFileName
}

func CheckAndPrintErrors(errors ...error) {
	for _, err := range errors {
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
}
