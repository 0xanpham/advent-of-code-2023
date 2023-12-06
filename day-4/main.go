package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	DOT        = "DOT"
	NUMBER     = "NUMBER"
	SYMBOL     = "SYMBOL"
	STAR       = "STAR"
)

func totalPoint(data string) int {
	winningNumbersStr := strings.Split(data, " | ")[0]
	numbersStr := strings.Split(data, " | ")[1]
	if winningNumbersStr[0] == ' ' {
		winningNumbersStr = strings.Replace(winningNumbersStr, " ", "", 1)
	}
	if numbersStr[0] == ' ' {
		numbersStr = strings.Replace(numbersStr, " ", "", 1)
	}
	winningNumbers := strings.Split(strings.Replace(winningNumbersStr, "  ", " ", -1), " ")
	numbers := strings.Split(strings.Replace(numbersStr, "  ", " ", -1), " ")
	winningNumberMap := make(map[string]int)
	for _, number := range winningNumbers {
		winningNumberMap[number] = 1
	}
	point := 0
	for _, number := range numbers {
		if winningNumberMap[number] == 1 {
			if point == 0 {
				point = 1
			} else {
				point *= 2
			}
		}
	}
	return point
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	rows := strings.Split(stringContent, "\n")
	sumPart1 := 0
	for _, row := range rows {
		sumPart1 += totalPoint(strings.Split(row, ": ")[1])
	}
	fmt.Printf("Part 1 Result: %v\n", sumPart1)
	// fmt.Printf("Part 2 Result: %v\n", sumPart2)
}