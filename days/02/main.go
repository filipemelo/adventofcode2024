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
	safeProblemDampener := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		rowNumbers := getRowNumbersFromParts(parts)

		if !firstPartUnsafe(rowNumbers) {
			safe += 1
		} else {
			if !secondPartUnsafe(rowNumbers) {
				safeProblemDampener += 1
			}
		}
	}

	fmt.Println("Safe reports:", safe)
	fmt.Println("Safe reports problem dampener:", (safe + safeProblemDampener))
}

func firstPartUnsafe(rowNumbers []int) bool {
	isDecreasing := true
	isIncreasing := true

	for i := 1; i < len(rowNumbers); i++ {
		previousLevel := rowNumbers[i-1]
		currentLevel := rowNumbers[i]
		abs := internal.AbsDiff(previousLevel, currentLevel)

		if abs < 1 || abs > 3 {
			return true
		}

		if currentLevel > previousLevel {
			isDecreasing = false
		}
		if currentLevel < previousLevel {
			isIncreasing = false
		}
	}

	return !(isIncreasing || isDecreasing)
}

func getRowNumbersFromParts(parts []string) []int {
	var row []int
	for i := 0; i < len(parts); i++ {
		row = append(row, internal.ReadInt(parts[i]))
	}
	return row
}

func secondPartUnsafe(rowNumbers []int) bool {
	for i := 0; i < len(rowNumbers); i++ {
		modifiedRow := removeElement(rowNumbers, i)

		if !firstPartUnsafe(modifiedRow) {
			return false
		}
	}
	return true
}

func removeElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
