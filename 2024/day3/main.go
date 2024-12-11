package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
	"strconv"
)


func main() {
	inputString, _ := utils.ReadFile("input.txt")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(inputString, -1)

	total := 0

	for _, match := range matches {
		total += performInstruction(match)
	}

	fmt.Println("Total:", total)
}


func performInstruction(instruction string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	numbers := re.FindAllString(instruction, -1)

	y, _ := strconv.Atoi(numbers[0])
	x, _ := strconv.Atoi(numbers[1])

	return y * x
}