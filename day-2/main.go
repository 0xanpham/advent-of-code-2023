package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isAvailable(row string) bool {
	splitResult := strings.Split(row, ": ")
	picks := splitResult[1]
	Max := make(map[string]int)
	for _, pick := range strings.Split(picks, "; ") {
		for _, item := range strings.Split(pick, ", ") {
			data := strings.Split(item, " ")
			quantity, _ := strconv.Atoi(data[0])
			color := data[1]
			if quantity > Max[color] {
				Max[color] = quantity
			}
		}
	}
	return Max["red"] <= 12 && Max["green"] <= 13 && Max["blue"] <= 14
}

func power(row string) int {
	splitResult := strings.Split(row, ": ")
	picks := splitResult[1]
	Max := make(map[string]int)
	for _, pick := range strings.Split(picks, "; ") {
		for _, item := range strings.Split(pick, ", ") {
			data := strings.Split(item, " ")
			quantity, _ := strconv.Atoi(data[0])
			color := data[1]
			if quantity > Max[color] {
				Max[color] = quantity
			}
		}
	}
	return Max["red"] * Max["green"] * Max["blue"]
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	rows := strings.Split(stringContent, "\n")
	sumPart1 := 0
	sumPart2 := 0
	for i, row := range rows {
		if isAvailable(row) {
			sumPart1 += i + 1
		}
		sumPart2 += power(row)
	}
	fmt.Printf("Part 1 Result: %v\n", sumPart1)
	fmt.Printf("Part 2 Result: %v\n", sumPart2)
}