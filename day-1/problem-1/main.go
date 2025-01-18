package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
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

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	var difference float64

	for i := 0; i < len(leftNumbers); i++ {
		diff := float64(leftNumbers[i] - rightNumbers[i])
		difference += math.Abs(diff)
	}

	fmt.Println(int(difference))

}
