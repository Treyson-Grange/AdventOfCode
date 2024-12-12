package utils

import (
	"fmt"
	"time"
)

// Function to time a function
// Usage: utils.TimeFunction(func() { part1(inputString) })
func TimeFunction(f func()) {
	start := time.Now()
	f()
	fmt.Println("Function Took:", time.Since(start))	
}