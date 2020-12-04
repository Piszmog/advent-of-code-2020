package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	passwordsFile := flag.String("f", "passwords.txt", "File containing the passwords")

	passwords, err := readPasswords(*passwordsFile)
	if err != nil {
		log.Fatalln(err)
	}

	// Part 1
	validPasswords1 := getValidPasswordsPart1(passwords)
	fmt.Printf("Part 1: There are %d valid passwords\n", len(validPasswords1))

	// Part 2
	validPasswords2 := getValidPasswordsPart2(passwords)
	fmt.Printf("Part 2: There are %d valid passwords\n", len(validPasswords2))
}

type password struct {
	min    int
	max    int
	letter string
	value  string
}

func readPasswords(expenseFile string) ([]password, error) {
	file, err := os.Open(expenseFile)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to open file %s: %w", expenseFile, err)
	}
	defer file.Close()

	var passwords []password
	scanner := bufio.NewScanner(file)
	rowNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		// get the min/max
		minMax := strings.Split(lineParts[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert min %s: %w", minMax[0], err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert max %s: %w", minMax[1], err)
		}
		// get the letter
		letter := lineParts[1][:1]
		// get the value
		value := lineParts[2]
		passwords = append(passwords, password{
			min:    min,
			max:    max,
			letter: letter,
			value:  value,
		})
		rowNumber++
	}
	return passwords, err
}

func getValidPasswordsPart1(passwords []password) []password {
	var validPasswords []password
	for _, password := range passwords {
		count := strings.Count(password.value, password.letter)
		if count >= password.min && count <= password.max {
			validPasswords = append(validPasswords, password)
		}
	}
	return validPasswords
}

func getValidPasswordsPart2(passwords []password) []password {
	var validPasswords []password
	for _, password := range passwords {
		value := password.value
		r := []rune(password.letter)[0]
		hasPosition1 := false
		hasPosition2 := false
		if len(value) >= password.min {
			hasPosition1 = rune(value[password.min-1]) == r
		}
		if len(value) >= password.max {
			hasPosition2 = rune(value[password.max-1]) == r
		}
		if (hasPosition1 && !hasPosition2) || (!hasPosition1 && hasPosition2) {
			validPasswords = append(validPasswords, password)
		}
	}
	return validPasswords
}
