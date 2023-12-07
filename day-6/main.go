package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	inputs := strings.Split(stringContent, "\n")

	times := strings.Split(strings.Split(inputs[0], ":        ")[1], "     ")
	distances := strings.Split(strings.Split(inputs[1], ":   ")[1], "   ")
	resultPart1 := 1
	for i := range times {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		num := 0
		for j := 0; j <= time; j += 1 {
			if j * (time - j) > distance {
				num += 1 
			}
		}
		resultPart1 *= num
	}
	fmt.Printf("Part 1 Result: %v\n", resultPart1)
}