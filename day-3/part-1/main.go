package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	// Regular expression to match valid mul(X,Y) patterns
	// where X and Y are 1-3 digit numbers
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	for scanner.Scan() {
		line := scanner.Text()

		// Find all matches in the line
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			// Convert strings to integers
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])

			// Multiply and add to sum
			result := x * y
			sum += result
		}
	}

	fmt.Printf("Sum of all multiplication results: %d\n", sum)
}
