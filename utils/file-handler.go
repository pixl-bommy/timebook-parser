package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load json file
func LoadJsonFile[T any](filePath string) (*T, error) {
	// Read file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var output T

	err = json.Unmarshal(fileContent, &output)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json file: %w", err)
	}

	return &output, nil
}

// Save json to file
func StoreJsonFile[T any](filePath string, content *T) error {
	// encode json content to bytes
	data, err := json.Marshal(content)
	if err != nil {
		return fmt.Errorf("failed to encode json content: %w", err)
	}

	// write file
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

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
