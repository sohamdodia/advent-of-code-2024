package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()
		if isSafeReport(parseLine(line)) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

// parseLine converts a space-separated string of numbers into a slice of integers
func parseLine(line string) []int {
	numbers := strings.Fields(line)
	result := make([]int, len(numbers))

	for i, num := range numbers {
		val, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		result[i] = val
	}

	return result
}

// isSafeReport checks if a report follows the safety rules
func isSafeReport(levels []int) bool {
	if len(levels) < 2 {
		return true
	}

	// Check first difference to determine if we should be increasing or decreasing
	isIncreasing := levels[1] > levels[0]

	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]

		// Check if difference is between 1 and 3 (inclusive)
		if abs(diff) < 1 || abs(diff) > 3 {
			return false
		}

		// Check if direction matches the initial direction
		if isIncreasing && diff <= 0 {
			return false
		}
		if !isIncreasing && diff >= 0 {
			return false
		}
	}

	return true
}

// abs returns the absolute value of an integer
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
