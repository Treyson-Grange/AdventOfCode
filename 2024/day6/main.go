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

	var matrix [][]string

	var guardLocation [2]int

	for _, line := range lines {
		var row []string
		for _, letter := range line {
			if letter == '^' {
				guardLocation = [2]int{len(matrix), len(row)}
			}
			row = append(row, string(letter))
		}
		matrix = append(matrix, row)
	}

	utils.TimeFunction(func() { part1(matrix, guardLocation) })
	utils.TimeFunction(func() { part2(matrix) })
}

func part1(matrix [][]string, guardLocation [2]int) {
    squaresTraversed := 0
    directions := map[string][2]int{
        "up":    {-1, 0},
        "down":  {1, 0},
        "left":  {0, -1},
        "right": {0, 1},
    }
    nextDirection := map[string]string{
        "up":    "right",
        "right": "down",
        "down":  "left",
        "left":  "up",
    }
    direction := "up"

    for {
        if isAtEdge(matrix, guardLocation[0], guardLocation[1]) {
            fmt.Println("Squares traversed:", squaresTraversed)
            break
        }

        delta := directions[direction]
        newRow, newCol := guardLocation[0]+delta[0], guardLocation[1]+delta[1]

        if matrix[newRow][newCol] == "#" {
            direction = nextDirection[direction]
        } else {
            if !checkTraversed(matrix, newRow, newCol) {
                squaresTraversed++
                matrix[newRow][newCol] = "X"
            }
            guardLocation = [2]int{newRow, newCol}
        }
    }
}

func part2(matrix [][]string) {
	fmt.Println("Part 2")
	// We will ahve to check for loops. 
	// one way we might want 
	// Damn im so adhd, i shoudnt be doing this rn
}


// Queries if the given coordinates are at the edge of the matrix, therefore ending the program
func isAtEdge(matrix [][]string, i int, j int) bool {
	return i == 0 || i == len(matrix) - 1 || j == 0 || j == len(matrix[i]) - 1
}

func checkTraversed(matrix [][]string, x, y int) bool {
	return matrix[x][y] == "X"	
}