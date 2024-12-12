package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/filipemelo/adventofcode2024/internal"
)

// https://adventofcode.com/2024/day/2

func main() {
	file := internal.ReadFile("input.txt", "02")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safe := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		unsafe := false
		allDecreasing := (internal.ReadInt(parts[1]) < internal.ReadInt(parts[0]))
		allIncreasing := !allDecreasing
		for i := 1; i < len(parts); i++ {
			previousLevel := internal.ReadInt(parts[i-1])
			currentLevel := internal.ReadInt(parts[i])
			abs := internal.AbsDiff(previousLevel, currentLevel)

			if abs < 1 || abs > 3 {
				unsafe = true
				break
			}

			if allDecreasing && (currentLevel > previousLevel) {
				unsafe = true
				break
			}

			if allIncreasing && (currentLevel < previousLevel) {
				unsafe = true
				break
			}
		}

		if !unsafe {
			safe += 1
		}
	}

	fmt.Println("Safe reports:", safe)
}
