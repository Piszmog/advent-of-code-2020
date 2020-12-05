package main

import (
	"flag"
	"fmt"
	"github.com/Piszmog/adventofcode/utils"
	"log"
	"sort"
	"strings"
)

type boardingPass struct {
	value  []string
	row    int
	column int
	seatId int
}

type boardingPasses []boardingPass

func (p boardingPasses) Len() int {
	return len(p)
}
func (p boardingPasses) Less(i, j int) bool {
	return p[i].seatId < p[j].seatId
}
func (p boardingPasses) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	inputFile := flag.String("f", "input.txt", "File containing the input")

	var passes boardingPasses
	err := utils.ReadTextFile(*inputFile, func(line string) error {
		passes = append(passes, boardingPass{value: strings.Split(line, "")})
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	highestSeatId := 0
	for i, pass := range passes {
		row, column := decodeValue(pass.value)
		seatId := getSeatId(row, column)
		pass.row = row
		pass.column = column
		pass.seatId = seatId
		if pass.seatId > highestSeatId {
			highestSeatId = pass.seatId
		}
		passes[i] = pass
	}
	fmt.Printf("Part 1: highest seat id: %d\n", highestSeatId)

	// part 2
	sort.Sort(passes)
	var missingSeatId int
	for i, pass := range passes {
		if i == len(passes)-1 {
			continue
		}
		// if the next seat is not +1, then +1 from current pass is the missing seat
		if pass.seatId != passes[i+1].seatId-1 {
			missingSeatId = pass.seatId + 1
		}
	}
	fmt.Printf("Part 2: Missing seat: %d\n", missingSeatId)
}

const (
	frontLowerHalf = "F"
	backUpperHalf  = "B"
	leftLowerHalf  = "L"
	rightUpperHalf = "R"

	minRow    = 0
	maxRow    = 127
	minColumn = 0
	maxColumn = 7
)

func decodeValue(value []string) (int, int) {
	currentMinRow := minRow
	currentMaxRow := maxRow
	currentMinColumn := minColumn
	currentMaxColumn := maxColumn
	for _, entry := range value {
		switch entry {
		case frontLowerHalf:
			currentMaxRow = ((currentMaxRow - currentMinRow) / 2) + currentMinRow
		case backUpperHalf:
			currentMinRow = (currentMaxRow+currentMinRow)/2 + 1
		case leftLowerHalf:
			currentMaxColumn = ((currentMaxColumn - currentMinColumn) / 2) + currentMinColumn
		case rightUpperHalf:
			currentMinColumn = (currentMaxColumn+currentMinColumn)/2 + 1
		}
	}
	return currentMinRow, currentMinColumn
}

func getSeatId(row int, column int) int {
	return row*8 + column
}
