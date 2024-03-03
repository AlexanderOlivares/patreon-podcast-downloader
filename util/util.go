package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

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
