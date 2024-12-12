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

func TestAbsDiff(t *testing.T) {
	num1 := 1
	num2 := 10
	expected1 := 9

	num3 := -2
	num4 := -12
	expected2 := 10

	absDiff1 := internal.AbsDiff(num1, num2)
	absDiff2 := internal.AbsDiff(num2, num1)
	absDiff3 := internal.AbsDiff(num3, num4)
	absDiff4 := internal.AbsDiff(num4, num3)

	assert.Equal(t, expected1, absDiff1)
	assert.Equal(t, expected1, absDiff2)
	assert.Equal(t, expected2, absDiff3)
	assert.Equal(t, expected2, absDiff4)
}

func TestReadInt(t *testing.T) {
	numStr1 := "-1"
	numStr2 := "0"
	numStr3 := "1"
	numStr4 := "-10000"
	numStr5 := "999999"

	assert.Equal(t, -1, internal.ReadInt(numStr1))
	assert.Equal(t, 0, internal.ReadInt(numStr2))
	assert.Equal(t, 1, internal.ReadInt(numStr3))
	assert.Equal(t, -10000, internal.ReadInt(numStr4))
	assert.Equal(t, 999999, internal.ReadInt(numStr5))
}
