package main

import (
	"flag"
	"fmt"
	"github.com/Piszmog/adventofcode/utils"
	"log"
	"strings"
)

type group struct {
	people []person
}

type person struct {
	yesQuestions []string
}

type void struct{}

var member void

func main() {
	inputFile := flag.String("f", "input.txt", "File containing the input")

	groups, err := getGroups(inputFile)
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	anyAnsweredQuestionsCount := 0
	for _, g := range groups {
		uniqueQuestions := make(map[string]void)
		for _, p := range g.people {
			for _, q := range p.yesQuestions {
				uniqueQuestions[q] = member
			}
		}
		anyAnsweredQuestionsCount += len(uniqueQuestions)
	}
	fmt.Printf("Part 1: Total answered questions (any): %d\n", anyAnsweredQuestionsCount)

	// part 2
	allAnsweredQuestionsCount := 0
	for _, g := range groups {
		groupSize := len(g.people)
		uniqueQuestions := make(map[string]int)
		for _, p := range g.people {
			for _, q := range p.yesQuestions {
				uniqueQuestions[q] += 1
			}
		}
		for _, count := range uniqueQuestions {
			if count == groupSize {
				allAnsweredQuestionsCount += 1
			}
		}
	}
	fmt.Printf("Part 2: Total answered questions (all): %d\n", allAnsweredQuestionsCount)
}

func getGroups(inputFile *string) ([]group, error) {
	var groups []group
	var currentGroup group
	err := utils.ReadTextFile(*inputFile, func(line string) error {
		if len(line) == 0 {
			newGroup := currentGroup
			groups = append(groups, newGroup)
			currentGroup.people = nil
			return nil
		}
		questions := strings.Split(line, "")
		currentGroup.people = append(currentGroup.people, person{yesQuestions: questions})
		return nil
	})
	if err != nil {
		return nil, err
	}
	groups = append(groups, currentGroup)
	return groups, nil
}
