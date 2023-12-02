package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var spellDigitReplacer *strings.Replacer

func inputFromFile(filename string) []byte {
	binp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return binp
}

func extractDigits(calString string) int {
	runes := []rune(calString)
	l_found, r_found := ' ', ' '

	for l := 0; l <= len(calString)-1; l++ {
		if unicode.IsNumber(runes[l]) {
			l_found = runes[l]
			break
		}
	}
	for r := len(calString)-1; r >= 0; r-- {
		if unicode.IsNumber(runes[r]) {
			r_found = runes[r]
			break
		}
	}

	digit := string(l_found) + string(r_found)
	digit = strings.TrimSpace(digit)

	intDigit, e := strconv.Atoi(digit)
	if e != nil {
		fmt.Println("Error Converting Inputs: '", e, "' Defaulting to 0")
		return 0
	}

	return intDigit
}

func convertSpellDigits(calString string) string {
	return spellDigitReplacer.Replace(calString)
}

func parseCalValues(calibValues []string, spellDigit bool) []int {
	var values []int
	for _, calString := range calibValues {

		// task 2
		if spellDigit {
			// super hacky! will probably convert it to regex based soln
			calString = convertSpellDigits(calString)
			calString = convertSpellDigits(calString)
		}
		
		values = append(values, extractDigits(calString))
	}
	return values
}

func calibSum(digits []int) int {
	var tot int
	for _, v := range digits {
		tot += v
	}

	return tot
}

func main() {
	spellDigitReplacer = strings.NewReplacer(
		"one", "1e",
		"two", "2o",
		"three", "3e",
		"four", "4r",
		"five", "5e",
		"six", "6x",
		"seven", "7n",
		"eight", "8t",
		"nine", "9e",
	)

	// Read input file
	inp := inputFromFile("inputs/input")

	// convert to slice of strings
	calibStrings := strings.Split(string(inp), "\n")

	// task 1
	// iterate over the data and extract first and last digit
	calibDigits := parseCalValues(calibStrings, false)
	task1 := calibSum(calibDigits)

	fmt.Println("Task 1:", task1)

	// task 2
	calibDigits = parseCalValues(calibStrings, true)
	task2 := calibSum(calibDigits)

	fmt.Println("Task 2:", task2)
}
