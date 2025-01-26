package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var allMatches [][]string
	for scanner.Scan() {
		allMatches = append(allMatches, findValidMulsDoDonts(scanner.Text())...)
	}

	fmt.Println(findAnswer(allMatches))
}

func findValidMulsDoDonts(input string) [][]string {
	r := regexp.MustCompile(`(mul\([\d]{1,3},[0-9]{1,3}\))|don't\(\)|do\(\)`)
	return r.FindAllStringSubmatch(input, -1)
}

func findProductInPattern(input string) int {
	r := regexp.MustCompile(`\d+`)
	nums := r.FindAllString(input, -1)
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1 * n2
}

func findAnswer(input [][]string) int {
	sum := 0
	active := true
	for _, match := range input {
		switch {
		case strings.Contains(match[0], "don"):
			active = false
		case strings.Contains(match[0], "do("):
			active = true
		default:
			if active {
				sum += findProductInPattern(match[0])
			}
		}
	}
	return sum
}
