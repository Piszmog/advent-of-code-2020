package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	treeCharacter = "#"
)

type slope struct {
	down  int
	right int
}

func main() {
	terrainFile := flag.String("f", "terrain.txt", "File containing the terrain")

	terrain, err := readTerrain(*terrainFile)
	if err != nil {
		log.Fatalln(err)
	}

	// Part 1
	encounter1 := getEncounters(terrain, slope{down: 1, right: 3})
	fmt.Printf("Part 1: Tree Squares: %d\n", encounter1.trees)

	// Part 2
	encounter2 := getEncounters(terrain, slope{down: 1, right: 1})
	encounter3 := getEncounters(terrain, slope{down: 1, right: 5})
	encounter4 := getEncounters(terrain, slope{down: 1, right: 7})
	encounter5 := getEncounters(terrain, slope{down: 2, right: 1})
	fmt.Printf("Part 2: Result: %d\n",
		encounter1.trees*encounter2.trees*encounter3.trees*encounter4.trees*encounter5.trees)
}

func readTerrain(expenseFile string) ([][]string, error) {
	file, err := os.Open(expenseFile)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file %s: %w", expenseFile, err)
	}
	defer file.Close()

	var terrain [][]string
	scanner := bufio.NewScanner(file)
	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		terrain = append(terrain, strings.Split(line, ""))
		rowNumber++
	}
	return terrain, err
}

type encounter struct {
	open  int
	trees int
}

func getEncounters(terrain [][]string, slope slope) encounter {
	row := 0
	column := 0
	openSquares := 0
	treeSquares := 0
	for row+slope.down < len(terrain) {
		row += slope.down
		nextTerrain := terrain[row]
		column = (column + slope.right) % len(nextTerrain)
		if nextTerrain[column] == treeCharacter {
			treeSquares++
		} else {
			openSquares++
		}
	}
	return encounter{open: openSquares, trees: treeSquares}
}
