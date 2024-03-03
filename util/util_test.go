package util

import (
	"testing"
)

func TestSanitizeFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"   example file.txt  ", "example_file.txt"},                                                  // Spaces trimmed and replaced with underscores
		{"file_with_special_!@#$%^*()characters.txt", "file_with_special_characters.txt"},              // Aside from &,:- special characters removed
		{"filename with spaces and special characters", "filename_with_spaces_and_special_characters"}, // Spaces and special characters removed
		{"file.with.dots.txt", "file.with.dots.txt"},                                                   // Dots retained
		{"file-with-hyphens.txt", "file-with-hyphens.txt"},                                             // Hyphens are ok
		{"file_underscored.txt", "file_underscored.txt"},                                               // Underscores retained
		{"", ""}, // Empty string remains empty
	}

	for _, test := range tests {
		result := SanitizeFileName(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
