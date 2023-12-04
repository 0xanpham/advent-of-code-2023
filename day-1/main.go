package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var DIGITS []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var value map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func getCalibrationValueV1(s string) int {
	var left int
	var right int
	runes := []rune(s)
	for i := 0; i < len(runes); i += 1 {
		if int(runes[i]) >= 48 && int(runes[i]) <= 57 {
			left = int(runes[i]) - 48
			break
		}
	}
	for i := len(runes) - 1; i >= 0; i -= 1 {
		if int(runes[i]) >= 48 && int(runes[i]) <= 57 {
			right = int(runes[i]) - 48
			break
		}
	}
	return left*10 + right
}

func getCalibrationValueV2(s string) int {
	var left int
	var right int
	leftIndex := math.MaxInt
	rightIndex := math.MinInt

	for _, digit := range DIGITS {
		index := strings.Index(s, digit)
		if index != -1 && index < leftIndex {
			leftIndex = index
			left = value[digit]
		}
	}

	for _, digit := range DIGITS {
		index := strings.LastIndex(s, digit)
		if index != -1 && index > rightIndex {
			rightIndex = index
			right = value[digit]
		}
	}

	runes := []rune(s)

	if leftIndex == math.MaxInt {
		leftIndex = len(runes)
	}
	if rightIndex == math.MinInt {
		rightIndex = 0
	}
	fmt.Printf("%v %v\n", leftIndex, rightIndex)

	for i := 0; i < leftIndex; i += 1 {
		if int(runes[i]) >= 48 && int(runes[i]) <= 57 {
			left = int(runes[i]) - 48
			break
		}
	}

	for i := len(runes) - 1; i >= rightIndex; i -= 1 {
		if int(runes[i]) >= 48 && int(runes[i]) <= 57 {
			right = int(runes[i]) - 48
			break
		}
	}
	fmt.Printf("%v%v\n", left, right)
	return left*10 + right
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
	for _, row := range rows {
		sumPart1 += getCalibrationValueV1(row)
		sumPart2 += getCalibrationValueV2(row)
	}
	fmt.Printf("Part 1 Answer: %v\n", sumPart1)
	fmt.Printf("Part 2 Answer: %v\n", sumPart2)
}
