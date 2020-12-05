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

const (
	birthYearField      = "byr"
	issueYearField      = "iyr"
	expirationYearField = "eyr"
	heightField         = "hgt"
	hairColorField      = "hcl"
	eyeColorField       = "ecl"
	passportIdField     = "pid"
	countryIdField      = "cid"
)

func main() {
	passportsFile := flag.String("f", "passports.txt", "File containing the terrain")

	passports, err := readPassports(passportsFile)
	if err != nil {
		log.Fatalln(err)
	}

	// part 1
	validPassports1 := getValidPassportsFieldExistence(passports)
	fmt.Printf("Part 1: Valid Passports: %d\n", len(validPassports1))

	// part 2
	validPassports2 := getValidPassportsValues(validPassports1)
	fmt.Printf("Part 2: Valid Passports: %d\n", len(validPassports2))
}

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	id             string
	countryId      string
}

func readPassports(passportsFile *string) ([]passport, error) {
	var passports []passport
	currentPassport := passport{}
	err := utils.ReadTextFile(*passportsFile, func(line string, rowNumber int) error {
		// on to a new record, save the passport off to the slice
		if len(line) == 0 {
			passports = append(passports, clonePassport(currentPassport))
			resetPassport(&currentPassport)
		}
		// else populate the field[s]
		for _, field := range strings.Split(line, " ") {
			fieldParts := strings.Split(field, ":")
			switch fieldParts[0] {
			case birthYearField:
				currentPassport.birthYear = fieldParts[1]
			case issueYearField:
				currentPassport.issueYear = fieldParts[1]
			case expirationYearField:
				currentPassport.expirationYear = fieldParts[1]
			case heightField:
				currentPassport.height = fieldParts[1]
			case hairColorField:
				currentPassport.hairColor = fieldParts[1]
			case eyeColorField:
				currentPassport.eyeColor = fieldParts[1]
			case passportIdField:
				currentPassport.id = fieldParts[1]
			case countryIdField:
				currentPassport.countryId = fieldParts[1]
			}
		}
		return nil
	})
	// make sure to process the last record ( :D )
	passports = append(passports, clonePassport(currentPassport))
	resetPassport(&currentPassport)
	return passports, err
}

func clonePassport(original passport) passport {
	return passport{
		birthYear:      original.birthYear,
		issueYear:      original.issueYear,
		expirationYear: original.expirationYear,
		height:         original.height,
		hairColor:      original.hairColor,
		eyeColor:       original.eyeColor,
		id:             original.id,
		countryId:      original.countryId,
	}
}

func resetPassport(p *passport) {
	p.birthYear = ""
	p.issueYear = ""
	p.expirationYear = ""
	p.height = ""
	p.hairColor = ""
	p.eyeColor = ""
	p.id = ""
	p.countryId = ""
}

func getValidPassportsFieldExistence(passports []passport) []passport {
	var validPassports []passport
	for _, passport := range passports {
		if hasRequiredFields(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

func hasRequiredFields(p passport) bool {
	return len(p.birthYear) != 0 && len(p.issueYear) != 0 && len(p.expirationYear) != 0 && len(p.height) != 0 &&
		len(p.hairColor) != 0 && len(p.eyeColor) != 0 && len(p.id) != 0
}

func getValidPassportsValues(passports []passport) []passport {
	var validPassports []passport
	for _, passport := range passports {
		if hasValidFields(passport) {
			validPassports = append(validPassports, passport)
		}
	}
	return validPassports
}

var birthYearPattern = regexp.MustCompile("^[0-9]{4}$")
var issueYearPattern = regexp.MustCompile("^[0-9]{4}$")
var expirationYearPattern = regexp.MustCompile("^[0-9]{4}$")
var heightPattern = regexp.MustCompile("^([0-9]{2,3})((cm)|(in))$")
var hairColorPattern = regexp.MustCompile("#[0-9a-f]{6}")
var eyeColorPattern = regexp.MustCompile("^(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)$")
var idPattern = regexp.MustCompile("^[0-9]{9}$")

func hasValidFields(p passport) bool {
	if !birthYearPattern.MatchString(p.birthYear) || !issueYearPattern.MatchString(p.issueYear) ||
		!expirationYearPattern.MatchString(p.expirationYear) || !heightPattern.MatchString(p.height) ||
		!hairColorPattern.MatchString(p.hairColor) || !eyeColorPattern.MatchString(p.eyeColor) ||
		!idPattern.MatchString(p.id) {
		return false
	}
	if year, _ := strconv.Atoi(p.birthYear); year < 1920 || year > 2002 {
		return false
	}
	if year, _ := strconv.Atoi(p.issueYear); year < 2010 || year > 2020 {
		return false
	}
	if year, _ := strconv.Atoi(p.expirationYear); year < 2020 || year > 2030 {
		return false
	}
	heightSubmatches := heightPattern.FindStringSubmatch(p.height)
	if heightSubmatches[2] == "in" {
		if height, _ := strconv.Atoi(heightSubmatches[1]); height < 59 || height > 76 {
			return false
		}
	} else if height, _ := strconv.Atoi(heightSubmatches[1]); height < 150 || height > 193 {
		return false
	}
	return true
}
