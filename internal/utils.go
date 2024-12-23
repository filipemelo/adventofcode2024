package internal

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func ReadFileAsString(filename string, dayStr string) string {
	_, run, _, _ := runtime.Caller(0)
	dir := filepath.Dir(run)

	filePath := filepath.Join(dir, "..", "days", dayStr, filename)
	fmt.Println(filePath)

	data, err := os.ReadFile(filePath)
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

func AbsDiff(num1 int, num2 int) int {
	return int(math.Abs(float64(num1 - num2)))
}

func ReadInt(numStr string) int {
	num, _ := strconv.Atoi(numStr)
	return num
}
