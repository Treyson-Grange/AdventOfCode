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
	fmt.Println("Part 1")
	fmt.Println(lines)

	rules := make(map[string]string)
	//91|27
// 74|64
// 74|99
// 34|26
	for _, line := range lines {
		items := strings.Split(line, "|")
		if len(items) != 2 {
			break;
		}
		rules[items[0]] = items[1]
	}
	fmt.Println(len(rules))
}