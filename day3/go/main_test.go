package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetEncounters(t *testing.T) {
	terrainString := `..#..#......###.#...#......#..#
...#.....#...#...#..........#..
....#.#...............#.#.#....
.........#.......##............
#.#....#.#####.##.#........#..#`

	terrain := convertTerrainString(terrainString)

	type test struct {
		slope         slope
		expectedTrees int
	}

	tests := []test{
		{slope: slope{down: 1, right: 1}, expectedTrees: 0},
		{slope: slope{down: 1, right: 2}, expectedTrees: 1},
		{slope: slope{down: 1, right: 3}, expectedTrees: 4},
		{slope: slope{down: 1, right: 4}, expectedTrees: 1},
		{slope: slope{down: 1, right: 5}, expectedTrees: 0},
		{slope: slope{down: 1, right: 6}, expectedTrees: 1},
		{slope: slope{down: 1, right: 7}, expectedTrees: 0},
		{slope: slope{down: 1, right: 7}, expectedTrees: 0},
		{slope: slope{down: 2, right: 1}, expectedTrees: 1},
		{slope: slope{down: 2, right: 3}, expectedTrees: 0},
	}

	for _, test := range tests {
		encounters := getEncounters(terrain, test.slope)
		assert.Equal(t, test.expectedTrees, encounters.trees,
			fmt.Sprintf("slope (%d, %d) failed", test.slope.down, test.slope.right))
	}
}

func convertTerrainString(input string) [][]string {
	var terrain [][]string
	row := strings.Split(input, "\n")
	for _, r := range row {
		terrain = append(terrain, strings.Split(r, ""))
	}
	return terrain
}
