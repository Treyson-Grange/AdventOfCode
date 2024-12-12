package main

import (
	"aoc2024/utils"
	"fmt"
	"regexp"
)


func main() {
	inputString, _ := utils.ReadFile("input.txt")	

	fmt.Println("Part 1")
	utils.TimeFunction(func() { part1(inputString) })
	fmt.Println("Part 2")
	utils.TimeFunction(func() { part2(inputString) })
}

func part1(input string) {
	
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(input, -1)

	fmt.Println(len(matches))

	total := 0

	for _, match := range matches {
		total += performInstruction(match)
	}

	fmt.Println("Total:", total)
}
func part2(input string) {
	total := 0
	do := true
	regex := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|\bdo\(\)|don't\(\)`)
	matches := regex.FindAllString(input, -1)

	for _, match := range matches {
		if match == "do()" {
			do = true
		} else if match == "don't()" {
			do = false
		} else if do {
			result := performInstruction(match)
			total += result
		}
	}
	
	fmt.Println("Total:", total)
}

// Function to perform the instruction
func performInstruction(instruction string) int {
	re := regexp.MustCompile(`\d{1,3}`)
	numbers := re.FindAllString(instruction, -1)

	y := utils.StringToInt(numbers[0])
	x := utils.StringToInt(numbers[1])

	return y * x
}