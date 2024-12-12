package internal

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("File reading error %s: %v", filename, err)
	}
	return strings.TrimSpace(string(data))
}
