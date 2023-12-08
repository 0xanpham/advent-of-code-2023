package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func part1TotalSteps(instructions string, left map[string]string, right map[string]string) int {
	steps := 0
	current := "AAA"
	for i := 0; i < len(instructions); i = (i + 1) % len(instructions) {
		instruction := instructions[i]
		if current == "ZZZ" {
			return steps
		}
		if instruction == 'L' {
			current = left[current]
		} else {
			current = right[current]
		}
		steps += 1
	}
	return -1
}

func totalSteps(instructions string, start string, left map[string]string, right map[string]string) int {
	steps := 0
	current := start
	for i := 0; i < len(instructions); i = (i + 1) % len(instructions) {
		instruction := instructions[i]
		if current[len(current)-1] == 'Z' {
			return steps
		}
		if instruction == 'L' {
			current = left[current]
		} else {
			current = right[current]
		}
		steps += 1
	}
	return -1
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a int, b int) int {
	return (a / gcd(a, b)) * b
}

func isEnding(currents []string) bool {
	for _, current := range currents {
		if current[len(current)-1] != 'Z' {
			return false
		}
	}
	return true
}

func part2TotalSteps(instructions string, starts []string, left map[string]string, right map[string]string) int {
	result := 1
	for _, start := range starts {
		result = lcm(result, totalSteps(instructions, start, left, right))
	}
	return result
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	inputs := strings.Split(stringContent, "\n\n")

	instructions := inputs[0]
	left := make(map[string]string)
	right := make(map[string]string)
	starts := []string{}
	for _, row := range strings.Split(inputs[1], "\n") {
		src := strings.Split(row, " = ")[0]
		if src[len(src)-1] == 'A' {
			starts = append(starts, src)
		}
		pair := strings.Split(row, " = ")[1]
		pair = strings.Replace(pair, "(", "", -1)
		pair = strings.Replace(pair, ")", "", -1)
		leftDest := strings.Split(pair, ", ")[0]
		rightDest := strings.Split(pair, ", ")[1]
		left[src] = leftDest
		right[src] = rightDest
	}
	fmt.Printf("Part 1 Result: %v\n", part1TotalSteps(instructions, left, right))
	fmt.Printf("Part 2 Result: %v\n", part2TotalSteps(instructions, starts, left, right))
}
