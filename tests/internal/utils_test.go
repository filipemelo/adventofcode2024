package internal_test

import (
	"testing"

	"github.com/filipemelo/adventofcode2024/internal"
	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	filename := "testdata/example.txt"
	expected := "This is just an exemple text file =)"

	actual := internal.ReadFile(filename)
	assert.Equal(t, expected, actual, "The contented doesn't match.")
}
