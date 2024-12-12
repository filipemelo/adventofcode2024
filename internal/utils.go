package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func ReadFileAsString(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("File reading error %s: %v", filename, err)
	}
	return strings.TrimSpace(string(data))
}

func ReadFile(filename string, dayStr string) *os.File {
	_, run, _, _ := runtime.Caller(0)
	dir := filepath.Dir(run)

	filePath := filepath.Join(dir, "..", "days", dayStr, filename)
	fmt.Println(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error to open file: %v\n", err)

	}
	return file
}
