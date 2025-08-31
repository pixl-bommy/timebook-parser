package utils

import (
	"fmt"
	"os"
)

// Load file to string array
// Returns an array of strings of each line in the file, separated by newline
// or an error if the file could not be read.
// Empty lines are ignored.
func LoadFileToStringArray(filePath string) ([]string, error) {
	// Read file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	lines := RetrieveLinesFromContent(fileContent)

	return lines, nil
}

// Split file content into lines
func RetrieveLinesFromContent(fileContent []byte) []string {
	lines := make([]string, 0)

	currentLine := ""
	for _, char := range fileContent {
		switch char {
		case '\n':
			// Newline indicates end of line
			lines = append(lines, currentLine)
			currentLine = ""
			continue

		case '\r':
			// Ignore carriage return
			continue
		}

		// There are no "continues", so let's add character to current line
		currentLine += string(char)
	}

	// We've reached EOF, so add any remaining current line to lines array
	if len(currentLine) > 0 {
		lines = append(lines, currentLine)
	}

	return lines
}
