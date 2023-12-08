package common

import (
	"io"
	"jAPI/config"
	"log"
	"os"
	"strings"
)

func FileOrString(cfg *config.Config) ([]string, bool) {
	fileInfo, err := os.Stat(cfg.JOB)
	if err != nil {
		log.Println(err)
		return []string{}, false
	}

	if fileInfo.IsDir() {
		return nil, false
	}

	file, err := os.Open(cfg.JOB)
	if err != nil {
		log.Println(err)
		return nil, false
	}
	defer func() {
		_ = file.Close()
	}()

	fileData, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	fileLines := TrimString(string(fileData))

	if len(fileLines) == 1 && fileLines[0] == "" {
		return nil, false
	}

	return fileLines, true
}

func TrimString(str string) []string {
	elements := strings.Fields(str)

	for i, element := range elements {
		elements[i] = strings.TrimSpace(element)
	}
	return elements
}
