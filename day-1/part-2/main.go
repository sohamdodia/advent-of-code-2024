package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftNumbers []int
	var rightNumbers []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		leftVal, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		rightVal, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		leftNumbers = append(leftNumbers, leftVal)
		rightNumbers = append(rightNumbers, rightVal)
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		log.Fatal(err)
	}

	similarityScore := 0

	for i := 0; i < len(leftNumbers); i++ {
		count := 0
		for j := 0; j < len(rightNumbers); j++ {
			if leftNumbers[i] == rightNumbers[j] {
				count++
			}
		}
		similarityScore += count * leftNumbers[i]
	}

	fmt.Println(similarityScore)
}
