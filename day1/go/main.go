package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	expectedSum = 2020
)

func main() {
	expenseFile := flag.String("f", "../expenses.csv", "File containing the expenses")

	expenses, err := readExpenses(*expenseFile)
	if err != nil {
		log.Fatalln(err)
	}

	foundPart1 := false
	foundPart2 := false
	for i, expense1 := range expenses {
		for j, expense2 := range expenses {
			if i == j {
				continue
			}
			// for part 1, find 2 numbers that equal the expected sum
			if !foundPart1 && isExpectedPart1(expense1, expense2) {
				fmt.Printf("Part 1: Expense 1: %d, Expense 2: %d, Result: %d\n", expense1, expense2, expense1*expense2)
				foundPart1 = true
				// if we have found part 2 already, let's quit early
				if foundPart2 {
					return
				}
			}
			// for part 2, find 3 numbers that equal the expected sum
			for k, expense3 := range expenses {
				if i == j || i == k || j == k {
					continue
				}
				if !foundPart2 && isExpectedPart2(expense1, expense2, expense3) {
					fmt.Printf("Part 2: Expense 1: %d, Expense 2: %d, Expense 3: %d, Result: %d\n",
						expense1, expense2, expense3, expense1*expense2*expense3)
					foundPart2 = true
					// if we have found part 1 already, let's quit early
					if foundPart1 {
						return
					}
				}
			}
		}
	}
}

func readExpenses(expenseFile string) ([]int, error) {
	file, err := os.Open(expenseFile)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file %s: %w", expenseFile, err)
	}
	defer file.Close()

	var expenses []int
	reader := csv.NewReader(file)
	rowNumber := 0
	for {
		row, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return nil, fmt.Errorf("failed to read row %d: %w", rowNumber, err)
		}
		expense, err := strconv.Atoi(row[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert %s to int: %w", row[0], err)
		}
		expenses = append(expenses, expense)
		rowNumber++
	}
	return expenses, err
}

func isExpectedPart1(x int, y int) bool {
	return x+y == expectedSum
}

func isExpectedPart2(x int, y int, k int) bool {
	return x+y+k == expectedSum
}
