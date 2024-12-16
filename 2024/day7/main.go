package main

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func main() {
	lines, err := utils.ReadFileLines("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(lines)
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	// Basically, we just need to go through each line, and check if we can get the 
	// left side (answer) from the right side (a list of numbers).
	// EX: 190: 10 19 
	// This is possible, if we place a * between 10 and 19, we get 190
	// If we can find a combination that equals the answer, we add it to the total

	total := 0
	for _, line := range lines {
		sides := strings.Split(line, ": ")
		answer := utils.StringToInt(sides[0])
		numbers := strings.Split(sides[1], " ")
		if isSolvable(numbers, answer, 0) {
			total += answer
		}
	}
	fmt.Println(total)
}

func isSolvable(line []string, answer, currentSol int) bool {
	if currentSol == answer && len(line) == 0 {
		return true
	}
	if len(line) == 0 || currentSol > answer {
		return false
	}
	return isSolvable(line[1:], answer, currentSol + utils.StringToInt(line[0])) || isSolvable(line[1:], answer, currentSol * utils.StringToInt(line[0]))
}

func part2(lines []string) {
	total := 0
	for _, line := range lines {
		sides := strings.Split(line, ": ")
		answer := utils.StringToInt(sides[0])
		numbers := strings.Split(sides[1], " ")
		if isSolveable2(numbers, answer, 0) {
			total += answer
		}
	}
	fmt.Println(total)
}

func isSolveable2(line []string, answer, currentSol int) bool {
	// New operator  || combines the digits from its left and right side 
	// EX: 10 || 19 = 1019

	if currentSol == answer && len(line) == 0 {
		return true
	}
	if len(line) == 0 || currentSol > answer {
		return false
	}
	concatenated := utils.StringToInt(fmt.Sprintf("%d%d", currentSol, utils.StringToInt(line[0])))
	return isSolveable2(line[1:], answer, currentSol + utils.StringToInt(line[0])) || isSolveable2(line[1:], answer, currentSol * utils.StringToInt(line[0])) || isSolveable2(line[1:], answer, concatenated)
}