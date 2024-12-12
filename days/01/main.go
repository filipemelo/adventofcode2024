package main

import (
	"bufio"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/filipemelo/adventofcode2024/internal"
)

func main() {
	file := internal.ReadFile("input.txt", "01")
	defer file.Close()

	var col1, col2 []int
	dict := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Line bad formatted: %s\n", line)
			continue
		}

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			fmt.Printf("Error to convert to integer at line: %s\n", line)
			continue
		}

		col1 = append(col1, num1)
		col2 = append(col2, num2)
		if val, ok := dict[num2]; ok {
			dict[num2] = val + 1
		} else {
			dict[num2] = 1
		}
	}

	sort.Ints(col1)
	sort.Ints(col2)

	sum := 0
	sim := 0
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		mod := int(math.Abs(float64(diff)))
		sum += mod
		if val, ok := dict[col1[i]]; ok {
			sim += col1[i] * val
		}
	}

	fmt.Println("Sum diff: ", sum)
	fmt.Println("Similarity score: ", sim)
}
