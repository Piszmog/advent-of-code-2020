package main

import (
	"flag"
	"fmt"
	"github.com/Piszmog/adventofcode/utils"
	"log"
	"strconv"
)

func main() {
	inputFile := flag.String("f", "../input.txt", "File containing the input")

	var numbers []int
	err := utils.ReadTextFile(*inputFile, func(line string) error {
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	invalidNumber := getInvalidNumber(numbers, 25)
	fmt.Printf("Part 1: Invalid number: %d\n", invalidNumber)

	// part 2
	weakness := getWeakness(numbers, invalidNumber)
	fmt.Printf("Part 2: Weakness: %d\n", weakness)
}

func getInvalidNumber(numbers []int, preambleRange int) int {
	invalid := 0
	for i := preambleRange; i < len(numbers); i++ {
		n := numbers[i]
		if !hasSum(numbers, n, preambleRange, i) {
			invalid = n
			break
		}
	}
	return invalid
}

func hasSum(numbers []int, number int, preambleRange int, index int) bool {
	for i := index - preambleRange; i < index; i++ {
		for j := index - preambleRange; j < index; j++ {
			if i == j {
				continue
			}
			if numbers[i]+numbers[j] == number {
				return true
			}
		}
	}
	return false
}

func getWeakness(numbers []int, invalidNumber int) int {
	smallest := 0
	largest := 0
numLoop:
	for i := 0; i < len(numbers); i++ {
		runningTotal := numbers[i]
		smallest = numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			runningTotal += numbers[j]
			if numbers[j] > largest {
				largest = numbers[j]
			}
			if numbers[j] < smallest {
				smallest = numbers[j]
			}
			if runningTotal == invalidNumber {
				break numLoop
			} else if runningTotal > invalidNumber {
				break
			}
		}
	}
	return smallest + largest
}
