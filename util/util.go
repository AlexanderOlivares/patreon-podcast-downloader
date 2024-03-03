package util

import (
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
