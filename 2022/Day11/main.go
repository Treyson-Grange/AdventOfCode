package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Base AOC GO")
	fmt.Println("Should be setup to read in input file, txt")

	file := readInputFile("./input.txt") //File as string
	monkeys := monkeyParser(file)

	fmt.Println(monkeys)

	rounds(monkeys, 20)

}

type Monkey struct {
	Number         int
	Items          []int //should be able to add to this, so we should use slice
	Operation      string
	Test           string
	ThrowToTrue    int
	ThrowToFalse   int
	MonkeyInspects int
}

func rounds(monkeys []Monkey, numRounds int) {
	for i := 0; i < numRounds; i++ {
		for j := range monkeys {
			monkey := &monkeys[j]
			newItems := []int{}
			for _, item := range monkey.Items {
				item = doOp(monkey.Operation, item)
				monkey.MonkeyInspects++

				relief := item / 3
				operand := monkey.Test

				operandValue, err := strconv.Atoi(operand)
				if err != nil {
					log.Fatal(err)
				}

				if relief%operandValue == 0 {
					monkeys[monkey.ThrowToTrue].Items = append(monkeys[monkey.ThrowToTrue].Items, relief)
				} else {
					monkeys[monkey.ThrowToFalse].Items = append(monkeys[monkey.ThrowToFalse].Items, relief)
				}
			}
			monkey.Items = newItems
		}
	}
	max1, max2 := -1, -1
	for _, monkey := range monkeys {
		if monkey.MonkeyInspects > max1 {
			max2 = max1
			max1 = monkey.MonkeyInspects
		} else if monkey.MonkeyInspects > max2 {
			max2 = monkey.MonkeyInspects
		}
	}
	for _, monkey := range monkeys {
		fmt.Println(monkey.MonkeyInspects)
	}
	fmt.Printf("Answer to part 1 %v", max1*max2)
}

func doOp(operation string, item int) int {
	operationComponents := strings.Split(operation, " ")
	rightside := operationComponents[2:]
	left := rightside[0]
	operator := rightside[1]
	right := rightside[2]

	leftValue, err := strconv.Atoi(left)
	if err != nil {
		leftValue = item
	}

	rightValue, err := strconv.Atoi(right)
	if err != nil {
		rightValue = item
	}

	var result int
	switch operator {
	case "+":
		result = leftValue + rightValue
	case "*":
		result = leftValue * rightValue
	default:
		log.Fatalf("Unsupported operator: %s", operator)
	}
	return result
}

/**
* Parses the input file into a list of Monkey structs
 */
func monkeyParser(file string) []Monkey {
	monkeys := strings.Split(file, "\n\n")
	listOfMonkeys := make([]Monkey, 0)
	for i, monkey := range monkeys {
		m := Monkey{}
		for j, line := range strings.Split(monkey, "\n") {
			m.Number = i
			if j == 1 { // Get starting items.
				numbers := strings.Split(line, ": ")[1]
				items := strings.Split(numbers, ", ")

				convertedItems := make([]int, 0)
				for _, item := range items {
					num, err := strconv.Atoi(item)
					if err != nil {
						log.Fatal(err)
					}
					convertedItems = append(convertedItems, num)
				}
				m.Items = convertedItems
			} else if j == 2 { // Get operation
				operation := strings.Split(line, ": ")[1]
				m.Operation = operation
			} else if j == 3 { // Get test
				test := strings.Split(line, "by ")[1]
				m.Test = strings.TrimSpace(test)
			} else if j == 4 { // Get true
				trueThreshold := strings.Split(line, "monkey ")[1]
				trueNum, err := strconv.Atoi(trueThreshold)
				if err != nil {
					log.Fatal(err)
				}
				m.ThrowToTrue = trueNum
			} else if j == 5 { // Get false
				falseThreshold := strings.Split(line, "monkey ")[1]
				falseNum, err := strconv.Atoi(falseThreshold)
				if err != nil {
					log.Fatal(err)
				}
				m.ThrowToFalse = falseNum
			}
		}
		listOfMonkeys = append(listOfMonkeys, m)
	}
	return listOfMonkeys
}

/*
* Only reads in the input file, handling input errors, and returns the file contents as a string
 */
func readInputFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read in the file
	var sb strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sb.String()
}
