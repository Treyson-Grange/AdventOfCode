package main

import (
	"aoc2024/utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadFileLines("./input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Part 1")
	fmt.Println(lines)
}