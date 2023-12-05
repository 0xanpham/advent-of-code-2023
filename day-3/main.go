package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	DOT        = "DOT"
	NUMBER     = "NUMBER"
	SYMBOL     = "SYMBOL"
	STAR       = "STAR"
)

func typeOf(char byte) string {
	if int(char) == 46 {
		return DOT
	} else if int(char) >= 48 && int(char) <= 57 {
		return NUMBER
	} else if int(char) == 42 {
		return STAR
	} else {
		return SYMBOL
	}
}

func checkSurround(x int, y int, data [][]byte) [][2]int {
	xMax := len(data[y]) - 1
	yMax := len(data) - 1
	var result [][2]int
	if (x > 0 && y > 0 && typeOf(data[y-1][x-1]) == NUMBER) {
		result = append(result, [2]int{x-1, y-1}) 
	}
	if (y > 0 && typeOf(data[y-1][x]) == NUMBER) {
		result = append(result, [2]int{x, y-1})
	}
	if (y > 0 && x < xMax && typeOf(data[y-1][x+1]) == NUMBER) {
		result = append(result, [2]int{x+1, y-1})
	}
	if (x > 0 && typeOf(data[y][x-1]) == NUMBER) {
		result = append(result, [2]int{x-1, y})
	}
	if (x < xMax && typeOf(data[y][x+1]) == NUMBER) {
		result = append(result, [2]int{x+1, y})
	}
	if (x > 0 && y < yMax && typeOf(data[y+1][x-1]) == NUMBER) {
		result = append(result, [2]int{x-1, y+1})
	}
	if (y < yMax && typeOf(data[y+1][x]) == NUMBER) {
		result = append(result, [2]int{x, y+1})
	}
	if (x < xMax && y < yMax && typeOf(data[y+1][x+1]) == NUMBER) {
		result = append(result, [2]int{x+1, y+1})
	}
	return result
}

func getNumber(x int, y int, data [][]byte) int {
	if typeOf(data[y][x]) != NUMBER {
		return 0
	}
	left := x
	right := x
	for left >= 0 && typeOf(data[y][left]) == NUMBER {
		left -= 1
	}
	left += 1
	for right <= len(data[y]) - 1 && typeOf(data[y][right]) == NUMBER {
		right += 1
	}
	right -= 1
	var byteArr []byte
	for i := left; i <= right; i += 1 {
		byteArr = append(byteArr, data[y][i])
		data[y][i] = byte('.')
	}
	str := string(byteArr)
	result, _ := strconv.Atoi(str)
	return result
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	rows := strings.Split(stringContent, "\n")
	var dataPart1 [][]byte
	var dataPart2 [][]byte
	for _, row := range rows {
		var rowArr []byte
		var rowArr2 []byte
		for _, charStr := range strings.Split(row, "") {
			rowArr = append(rowArr, charStr[0])
			rowArr2 = append(rowArr2, charStr[0])
		}
		dataPart1 = append(dataPart1, rowArr)
		dataPart2 = append(dataPart2, rowArr2)
	}
	sumPart1 := 0
	for i := range dataPart1 {
		for j := range dataPart1[i] {
			if typeOf(dataPart1[i][j]) == SYMBOL || typeOf(dataPart1[i][j]) == STAR {
				surrounds := checkSurround(j, i, dataPart1)
				for _, surround := range surrounds {
					x := surround[0]
					y := surround[1]
					number := getNumber(x, y, dataPart1)
					sumPart1 += number
				}
			}
		}
	}
	sumPart2 := 0
	for i := range dataPart2 {
		for j := range dataPart2[i] {
			if typeOf(dataPart2[i][j]) == STAR {
				surrounds := checkSurround(j, i, dataPart2)
				var surroundNumbers []int
				for _, surround := range surrounds {
					x := surround[0]
					y := surround[1]
					number := getNumber(x, y, dataPart2)
					if number > 0 {
						surroundNumbers = append(surroundNumbers, number)
					}
				}
				if len(surroundNumbers) == 2 {
					sumPart2 += surroundNumbers[0] * surroundNumbers[1]
				}
			}
		}
	}
	fmt.Printf("Part 1 Result: %v\n", sumPart1)
	fmt.Printf("Part 2 Result: %v\n", sumPart2)
}