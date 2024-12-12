package internal_test

import (
	"bufio"
	"testing"

	"github.com/filipemelo/adventofcode2024/internal"
	"github.com/stretchr/testify/assert"
)

func TestReadFileAsString(t *testing.T) {
	filename := "testdata/example.txt"
	expected := "This is just an exemple text file =)"

	actual := internal.ReadFileAsString(filename)
	assert.Equal(t, expected, actual, "The contented doesn't match.")
}

func TestReadFile(t *testing.T) {
	filename := "input.txt"
	expectedFilePath := "adventofcode/days/01/input.txt"
	expectedFirstLine := "37033   48086"

	actualFile := internal.ReadFile(filename, "01")
	defer actualFile.Close()

	scanner := bufio.NewScanner(actualFile)
	scanner.Scan()
	firstLine := scanner.Text()

	assert.Contains(t, actualFile.Name(), expectedFilePath, "Filename doesn't match.")
	assert.Equal(t, expectedFirstLine, firstLine, "The contented doesn't match.")
}
