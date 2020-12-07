package main

import (
	"flag"
	"fmt"
	"github.com/Piszmog/adventofcode/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputFile := flag.String("f", "../input.txt", "File containing the input")

	bags, err := getBags(*inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	count := countBagContained(bags, "shiny gold")
	fmt.Printf("Part 1: Number of bags: %d\n", count)

	// part 2
	total := totalBags(bags, "shiny gold")
	fmt.Printf("Part 2: Number of bags: %d\n", total)
}

type bag struct {
	name        string
	allowedBags map[string]int
}

const (
	splitName            = "bags contain"
	splitAllowedBagsList = ","
)

func getBags(inputFile string) (map[string]bag, error) {
	bagDetailsRegex := regexp.MustCompile("([0-9]+)\\s(([\\w]+ )+)")
	bags := make(map[string]bag)
	err := utils.ReadTextFile(inputFile, func(line string) error {
		lineParts := strings.Split(line, splitName)
		allowedBags := make(map[string]int)
		if lineParts[1] != " no other bags." {
			bagList := strings.Split(lineParts[1], splitAllowedBagsList)
			for _, v := range bagList {
				details := bagDetailsRegex.FindStringSubmatch(v)
				number, _ := strconv.Atoi(details[1])
				allowedBags[strings.TrimSpace(details[2])] = number
			}
		}
		name := strings.TrimSpace(lineParts[0])
		bags[name] = bag{
			name:        name,
			allowedBags: allowedBags,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return bags, nil
}

func countBagContained(bags map[string]bag, targetName string) int {
	count := 0
	for name, bag := range bags {
		if name == targetName {
			continue
		}
		if hasTargetBag(bags, bag.allowedBags, targetName) {
			count++
		}
	}
	return count
}

func hasTargetBag(bags map[string]bag, allowedBags map[string]int, targetName string) bool {
	if allowedBags[targetName] > 0 {
		return true
	}
	for name := range allowedBags {
		if hasTargetBag(bags, bags[name].allowedBags, targetName) {
			return true
		}
	}
	return false
}

func totalBags(bags map[string]bag, target string) int {
	count := 0
	for name, allowed := range bags[target].allowedBags {
		count += allowed
		nestedBags := totalBags(bags, name)
		count += allowed * nestedBags
	}
	return count
}
