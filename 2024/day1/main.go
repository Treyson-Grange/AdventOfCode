package main

import (
	"fmt"
	"math"
	"strings"

	"aoc2024/utils"
)

func main() {
	lines, err := utils.ReadFileLines("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	utils.TimeFunction(func() { part1(lines) })

	utils.TimeFunction(func() { part2(lines) })
}

func part1(lines []string) {
	leftEntries := make([]int, 0)
	rightEntries := make([]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, "   ")

		left := utils.StringToInt(parts[0])
		right := utils.StringToInt(parts[1])
		
		leftEntries = append(leftEntries, left)
		rightEntries = append(rightEntries, right)
	}

	leftEntries = utils.SortInts(leftEntries)
	rightEntries = utils.SortInts(rightEntries)

	totalDiff := 0

	for i := 0; i < len(leftEntries); i++ {
		diff := math.Abs(float64(leftEntries[i] - rightEntries[i]))
		totalDiff += int(diff)
	}
	fmt.Println("Total difference:", totalDiff)
}

func part2(lines []string) {
	//how often each number from the left list appears in the right list

	//first calculate how many times each number appears in the right list
	rightList := make(map[int]int)

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		right:= utils.StringToInt(parts[1])
		rightList[right]++
	}

	// Then, iterate through the left list, checking the right list for each number

	similarity_score := 0

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		left := utils.StringToInt(parts[0])
		if count, ok := rightList[left]; ok {
			// We multiply the number of times the number appears in the right list by the number.
			// We add this to a total sum
			similarity_score += left * count
		}
	}

	fmt.Println("Similarity score:", similarity_score)
}

