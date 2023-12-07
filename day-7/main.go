package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	FIVE_OF_A_KIND = 7
	FOUR_OF_A_KIND = 6
	FULL_HOUSE = 5
	THREE_OF_A_KIND = 4
	TWO_PAIR = 3
	ONE_PAIR = 2
	HIGH_CARD = 1
)

func handType(hand string) int {
	totalChar := make(map[rune]int)
	for _, char := range hand {
		totalChar[char] += 1
	}
	firstMax := 0
	var charWithFirstMax rune
	for _, char := range []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'} {
		if totalChar[char] > firstMax {
			firstMax = totalChar[char]
			charWithFirstMax = char
		}
	}
	totalChar[charWithFirstMax] = 0
	secondMax := 0
	for _, char := range []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'} {
		if totalChar[char] > secondMax {
			secondMax = totalChar[char]
		}
	}
	if firstMax == 5 {
		return FIVE_OF_A_KIND
	}
	if firstMax == 4 {
		return FOUR_OF_A_KIND
	}
	if firstMax == 3 {
		if secondMax == 2 {
			return FULL_HOUSE
		}
		return THREE_OF_A_KIND
	}
	if firstMax == 2 {
		if secondMax == 2 {
			return TWO_PAIR
		}
		return ONE_PAIR
	}
	return HIGH_CARD
}

func compare(firstHand string, secondHand string) int {
	firstHandType := handType(firstHand)
	secondHandType := handType(secondHand)
	
	value := make(map[byte]int)
	value['A'] = 14
	value['K'] = 13
	value['Q'] = 12
	value['J'] = 11
	value['T'] = 10
	value['9'] = 9
	value['8'] = 8
	value['7'] = 7
	value['6'] = 6
	value['5'] = 5
	value['4'] = 4
	value['3'] = 3
	value['2'] = 2

	if firstHandType == secondHandType {
		for i := range firstHand {
			if value[firstHand[i]] != value[secondHand[i]] {
				return value[firstHand[i]] - value[secondHand[i]]
			}
		}
		return 0
	}
	return firstHandType - secondHandType
}

func insertionSort(handBidPairs [][2]string) {
	for i := 0; i < len(handBidPairs); i += 1 {
		pivot := handBidPairs[i]
		for j := i - 1; j >= 0; j -= 1 {
			if compare(handBidPairs[j][0], pivot[0]) < 0 {
				handBidPairs[j+1] = pivot
				break
			}
			temp := handBidPairs[j]
			handBidPairs[j] = handBidPairs[j+1]
			handBidPairs[j+1] = temp
		}
	}
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	inputs := strings.Split(stringContent, "\n")
	var handBidPairs [][2]string
	for _, input := range inputs {
		handBidPairs = append(handBidPairs, [2]string(strings.Split(input, " ")))
	}
	insertionSort(handBidPairs)
	resultPart1 := 0
	for i, pair := range handBidPairs {
		bid, _ := strconv.Atoi(pair[1])
		resultPart1 += (i + 1) * bid
	}
	fmt.Printf("Part 1 Result: %v\n", resultPart1)
}