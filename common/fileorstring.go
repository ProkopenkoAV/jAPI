package common

import (
	"fmt"
	"io"
	"jAPI/config"
	"os"
	"strings"
)

// FileOrString checks if a file exists and returns its content.
func FileOrString(cfg *config.Config) ([]string, bool, error) {
	fileInfo, err := os.Stat(cfg.JOB)
	if err != nil {
		return nil, false, fmt.Errorf("file not found: %s", cfg.JOB)
	}

	file, err := os.Open(fileInfo.Name())
	if err != nil {
		return nil, true, fmt.Errorf("could not open file: %s. Permission denied", cfg.JOB)
	}
	defer func() {
		_ = file.Close()
	}()

	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, true, fmt.Errorf("could not read file: %s", cfg.JOB)
	}
	if len(fileData) == 0 {
		return nil, true, fmt.Errorf("file is empty: %s", cfg.JOB)
	}
	fileLines := TrimString(string(fileData))

	return fileLines, true, nil
}

// TrimString trims spaces from a string.
func TrimString(str string) []string {
	elements := strings.Fields(str)

	for i, element := range elements {
		elements[i] = strings.TrimSpace(element)
	}

	return elements
}
