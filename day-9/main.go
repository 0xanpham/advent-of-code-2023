package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func reverseSlice(slice []int) []int {
	result := []int{}
	for i := len(slice) - 1; i >= 0; i -= 1 {
		result = append(result, slice[i])
	}
	return result
}

func extrapolatedValue(sequence []int) int {
	lastNums := []int{}
	l := len(sequence)
	for l >= 1{
		nextSequence := sequence
		isAllZero := nextSequence[0] == 0
		for i := 1; i < l; i += 1 {
			if nextSequence[i] != 0 {
				isAllZero = false
			}
			nextSequence[i-1] = nextSequence[i] - nextSequence[i-1]
		}
		lastNums = append(lastNums, nextSequence[l-1])
		if isAllZero {
			break
		}
		l -= 1
	}
	prev := 0
	for i := len(lastNums) - 2; i >= 0; i -= 1 {
		prev += lastNums[i]
	}
	return prev
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	inputs := strings.Split(stringContent, "\n")
	sum1 := 0
	sum2 := 0
	for _, input := range inputs {
		sequence := []int{}
		for _, str := range strings.Split(input, " ") {
			num, _ := strconv.Atoi(str)
			sequence = append(sequence, num)
		}
		reverseSequence := reverseSlice(sequence)
		sum1 += extrapolatedValue(sequence)
		sum2 += extrapolatedValue(reverseSequence)
	}
	fmt.Printf("Part 1 Result: %v\n", sum1)
	fmt.Printf("Part 2 Result: %v\n", sum2)
}
