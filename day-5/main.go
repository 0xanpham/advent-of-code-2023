package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func locationOfSeed(
	seed int,
	seedToSoilData []string,
	soilToFertilizerData []string,
	fertilizierToWaterData []string,
	waterToLightData []string,
	lightToTemperatureData []string,
	temperatureToHumidityData []string,
	humidityToLocationData []string,
) int {
	soil := findNext(seed, seedToSoilData)
	fertilizer := findNext(soil, soilToFertilizerData)
	water := findNext(fertilizer, fertilizierToWaterData)
	light := findNext(water, waterToLightData)
	temperature := findNext(light, lightToTemperatureData)
	humidity := findNext(temperature, temperatureToHumidityData)
	location := findNext(humidity, humidityToLocationData)
	return location
}

func findNext(source int, data []string) int {
	for i := 1; i < len(data); i += 1 {
		mapSeedToSoilInput := data[i]
		mapSeedToSoilData := strings.Split(mapSeedToSoilInput, " ")
		d, _ := strconv.Atoi(mapSeedToSoilData[0])
		s, _ := strconv.Atoi(mapSeedToSoilData[1])
		r, _ := strconv.Atoi(mapSeedToSoilData[2])
		if source >= s && source < s + r {
			return d + source - s
		}
	}
	return source
}

func main() {
	content, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	stringContent := string(content)
	inputs := strings.Split(stringContent, "\n\n")

	seedsInput := strings.Split(inputs[0], ": ")[1]
	seedStrings := strings.Split(seedsInput, " ")
	var seeds []int
	for _, seedString := range seedStrings {
		seed, _ := strconv.Atoi(seedString)
		seeds = append(seeds, seed)
	}

	seedToSoilData := strings.Split(inputs[1], "\n")
	soilToFertilizerData := strings.Split(inputs[2], "\n")
	fertilizierToWaterData := strings.Split(inputs[3], "\n")
	waterToLightData := strings.Split(inputs[4], "\n")
	lightToTemperatureData := strings.Split(inputs[5], "\n")
	temperatureToHumidityData := strings.Split(inputs[6], "\n")
	humidityToLocationData := strings.Split(inputs[7], "\n")

	resultPart1 := math.MaxInt
	resultPart2 := math.MaxInt
	for _, seed := range seeds {
		seedLocation := locationOfSeed(
			seed,
			seedToSoilData,
			soilToFertilizerData,
			fertilizierToWaterData,
			waterToLightData,
			lightToTemperatureData,
			temperatureToHumidityData,
			humidityToLocationData,
		)
		resultPart1 = int(math.Min(float64(resultPart1), float64(seedLocation)))
	}
	for i := 0; i < len(seeds); i += 2 {
		fmt.Println(seeds[i])
		for j := seeds[i]; j < seeds[i] + seeds[i+1]; j += 1 {
			seedLocation := locationOfSeed(
				j,
				seedToSoilData,
				soilToFertilizerData,
				fertilizierToWaterData,
				waterToLightData,
				lightToTemperatureData,
				temperatureToHumidityData,
				humidityToLocationData,
			)
			resultPart2 = int(math.Min(float64(resultPart2), float64(seedLocation)))
		}
	}
	fmt.Printf("Part 1 Result: %v\n", resultPart1)
	fmt.Printf("Part 2 Result: %v\n", resultPart2)
}