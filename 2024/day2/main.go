package main

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

func main() {
	// One report per line
	// each report has a lsit of numbers called levels. Sep by spaces
	// amount of levels is variable
	// We need to define safe or not, safe if the levels are increasing only or decreasing only,
	// And if any two adjacent levels differ by at least ONE and at most THREE
	// Count the number of safe reports

	lines, err := utils.ReadFileLines("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	safeReportsPart1 := 0
	safeReportsPart2 := 0
	// Parse the input
	for _, line := range lines {
		// Split the line into levels
		levels := strings.Split(line, " ")
		// Check if the levels are safe
        intLevels := make([]int, len(levels))
        for i, level := range levels {
            num := utils.StringToInt(level)
            intLevels[i] = num
        }
        if isSafe(intLevels) {
			safeReportsPart1++
		}
		if isSafe2(levels) {
			safeReportsPart2++
		}

	}
	fmt.Println("Safe reports:", safeReportsPart1)

	fmt.Println("Safe reports part 2:", safeReportsPart2)
}

// Horrible implementation, works tho
func isSafe(levels []int) bool {
    if len(levels) < 2 {
        return true // Single-level or empty reports are trivially safe
    }

    // Determine the direction: increasing or decreasing
    isIncreasing := levels[1] > levels[0]
    isValid := true

    for i := 1; i < len(levels); i++ {
        diff := levels[i] - levels[i-1]

        // Check the difference constraints
        if diff < -3 || diff > 3 || diff == 0 {
            isValid = false
            break
        }

        // Check consistency of direction
        if (diff > 0) != isIncreasing {
            isValid = false
            break
        }
    }

    return isValid
}
// isSafe2 checks if a report is safe, allowing for the removal of one level.
func isSafe2(levels []string) bool {
    // Convert levels from string to integers
    intLevels := make([]int, len(levels))
    for i, level := range levels {
        num := utils.StringToInt(level)
        intLevels[i] = num
    }

    // Check if the original report is safe
    if isSafe(intLevels) {
        return true
    }

    // Check if removing one level makes the report safe
    for i := range intLevels {
        // Create a new slice with one level removed
        trimmed := append(intLevels[:i], intLevels[i+1:]...)
        if isSafe(trimmed) {
            return true
        }
    }

    // Report is not safe
    return false
}