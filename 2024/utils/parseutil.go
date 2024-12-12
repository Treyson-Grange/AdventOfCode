package utils

import (
	"fmt"
	"strconv"
)

// Function to convert a string to an integer, handles error, so we don't repeat code.
func StringToInt(input string) int {
	output, error := strconv.Atoi(input)
	if error != nil {
		fmt.Println("Error converting string to int:", error)
		return 0
	}
	return output
}