package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/filipemelo/adventofcode2024/internal"
)

func main() {
	input := internal.ReadFileAsString("input.txt", "03")
	result1 := sumMul(input)
	println("[1] - Sum mul:", result1)

	parts := sliceDoDont(input)
	result2 := sumMul(parts)

	println("[2] - Sum mul:", result2)
}

func sumMul(part string) int {
	pattern := `mul\((\d+),(\d+)\)`
	regex, _ := regexp.Compile(pattern)

	result := 0
	matches := regex.FindAllStringSubmatch(part, -1)

	for _, m := range matches {
		//println(m[0], m[1], m[2])
		num1, _ := strconv.Atoi(m[1])
		num2, _ := strconv.Atoi(m[2])

		result += num1 * num2
	}

	return result
}

func sliceDoDont(s string) string {
	var builder strings.Builder
	do := true

	for i, c := range s {
		if do {
			if i > 7 {
				sub := s[i-7 : i]
				if sub == "don't()" {
					do = false
				}
			}

			builder.WriteRune(c)
		} else {
			if i > 4 {
				sub := s[i-4 : i]
				if sub == "do()" {
					do = true

					builder.WriteRune(c)
				}
			}
		}
	}

	return builder.String()
}
