package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func getCalibrationValue(s string) int {
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
	return left * 10 + right
}

func main() {
	content, err := os.ReadFile("input.txt")

     if err != nil {
          log.Fatal(err)
     }

		stringContent := string(content)
		rows := strings.Split(stringContent, "\n")
		sum := 0
    for _, row := range rows {
			sum += getCalibrationValue(row)
		}
		fmt.Printf("Answer: %v\n", sum)
}