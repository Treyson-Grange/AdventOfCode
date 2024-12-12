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

	for _, line := range lines {
		var row []string
		for _, letter := range line {
			row = append(row, string(letter))
		}
		matrix = append(matrix, row)
	}

	fmt.Println("Part 1")
	utils.TimeFunction(func() { part1(matrix) })
	fmt.Println("Part 2")
	utils.TimeFunction(func() { part2(matrix) })
}


func part1(matrix [][]string) {
	// We have a matrix, we need to go through each row and column, and if we see an X,
	// we need to check all directions to see if there is an M, and if there is, we need to
	// Go the same direction, and see if there is an A, and if there is, we need to go the same
	// direction, and see if there is an S. If there is, we increment a total
	totalXMAS := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "X" {
				// Up
				if i-1 >= 0 && matrix[i-1][j] == "M" {
					if i-2 >= 0 && matrix[i-2][j] == "A" {
						if i-3 >= 0 && matrix[i-3][j] == "S" {
							totalXMAS++
						}
					}
				}
				// Down
				if i+1 < len(matrix) && matrix[i+1][j] == "M" {
					if i+2 < len(matrix) && matrix[i+2][j] == "A" {
						if i+3 < len(matrix) && matrix[i+3][j] == "S" {
							totalXMAS++
						}
					}
				}
				// Left
				if j-1 >= 0 && matrix[i][j-1] == "M" {
					if j-2 >= 0 && matrix[i][j-2] == "A" {
						if j-3 >= 0 && matrix[i][j-3] == "S" {
							totalXMAS++
						}
					}
				}
				// Right
				if j+1 < len(matrix[i]) && matrix[i][j+1] == "M" {
					if j+2 < len(matrix[i]) && matrix[i][j+2] == "A" {
						if j+3 < len(matrix[i]) && matrix[i][j+3] == "S" {
							totalXMAS++
						}
					}
				}
				// Up Left
				if i-1 >= 0 && j-1 >= 0 && matrix[i-1][j-1] == "M" {
					if i-2 >= 0 && j-2 >= 0 && matrix[i-2][j-2] == "A" {
						if i-3 >= 0 && j-3 >= 0 && matrix[i-3][j-3] == "S" {
							totalXMAS++
						}
					}
				}
				// Up Right
				if i-1 >= 0 && j+1 < len(matrix[i]) && matrix[i-1][j+1] == "M" {
					if i-2 >= 0 && j+2 < len(matrix[i]) && matrix[i-2][j+2] == "A" {
						if i-3 >= 0 && j+3 < len(matrix[i]) && matrix[i-3][j+3] == "S" {
							totalXMAS++
						}
					}
				}
				// Down Left
				if i+1 < len(matrix) && j-1 >= 0 && matrix[i+1][j-1] == "M" {
					if i+2 < len(matrix) && j-2 >= 0 && matrix[i+2][j-2] == "A" {
						if i+3 < len(matrix) && j-3 >= 0 && matrix[i+3][j-3] == "S" {
							totalXMAS++
						}
					}
				}
				// Down Right
				if i+1 < len(matrix) && j+1 < len(matrix[i]) && matrix[i+1][j+1] == "M" {
					if i+2 < len(matrix) && j+2 < len(matrix[i]) && matrix[i+2][j+2] == "A" {
						if i+3 < len(matrix) && j+3 < len(matrix[i]) && matrix[i+3][j+3] == "S" {
							totalXMAS++
						}
					}
				}
			}
		}
	}
	fmt.Println("Total XMAS:", totalXMAS)
}


func part2(matrix [][]string) {
	totalXMas := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "A" {
				upLeft := ""
				if i-1 >= 0 && j-1 >= 0 {
					upLeft = matrix[i-1][j-1]
				}
				bottomRight := ""
				if i+1 < len(matrix) && j+1 < len(matrix[i]) {
					bottomRight = matrix[i+1][j+1]
				}

				upRight := ""
				if i-1 >= 0 && j+1 < len(matrix[i]) {
					upRight = matrix[i-1][j+1]
				}
				bottomLeft := ""
				if i+1 < len(matrix) && j-1 >= 0 {
					bottomLeft = matrix[i+1][j-1]
				}

				// Check if the diagonals match the "XMas" pattern:
				if (upLeft == "M" && bottomRight == "S") || (upLeft == "S" && bottomRight == "M") {
					if (upRight == "M" && bottomLeft == "S") || (upRight == "S" && bottomLeft == "M") {
						totalXMas++
					}
				}
			}
		}
	}

	fmt.Println("Total XMas:", totalXMas)
}