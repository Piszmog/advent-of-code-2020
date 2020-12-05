package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetValidPassports(t *testing.T) {
	passports := []map[string]string{
		{
			"byr": "test",
			"iyr": "test",
			"eyr": "test",
			"hgt": "test",
			"hcl": "test",
			"ecl": "test",
			"pid": "test",
			"cid": "test",
		},
		{
			"byr": "test",
			"iyr": "test",
			"eyr": "test",
			"hcl": "test",
			"ecl": "test",
			"pid": "test",
			"cid": "test",
		},
		{
			"byr": "test",
			"iyr": "test",
			"eyr": "test",
			"hgt": "test",
			"hcl": "test",
			"ecl": "test",
			"pid": "test",
		},
		{
			"iyr": "test",
			"eyr": "test",
			"hgt": "test",
			"hcl": "test",
			"ecl": "test",
			"pid": "test",
		},
	}

	validPassports := getValidPassportsFieldExistence(passports, birthYearField, issueYearField, expirationYearField, heightField,
		hairColorField, eyeColorField, passportIdField)
	assert.Equal(t, 2, len(validPassports))
}
