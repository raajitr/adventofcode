package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func inputFromFile(filename string) []byte {
	binp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return binp
}

func extractDigits(calString string) int {
	l, r := 0, len(calString)-1
	runes := []rune(calString)
	l_found, r_found := ' ', ' '

	for l <= r {
		if unicode.IsNumber(runes[l]) && l_found == ' ' {
			l_found = runes[l]
		} else if unicode.IsNumber(runes[r]) && r_found == ' ' {
			r_found = runes[r]
		} else if l_found == ' ' {
			l += 1
		} else if r_found == ' ' {
			r -= 1
		}

		if l_found != ' ' && r_found != ' ' {
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

func parseCalValues(calibValues []string) []int {
	var values []int
	for _, calString := range calibValues {
		values = append(values, extractDigits(calString))
	}
	return values
}

func calibSum(digits []int) int {
	var tot int
	for _, v := range(digits) {
		tot += v
	}

	return tot
}


func main() {
	// Read input file
	inp := inputFromFile("puzzle_input.txt")

	// convert to slice of strings
	calibStrings := strings.Split(string(inp), "\n")

	// iterate over the data and extract first and last digit
	calibDigits := parseCalValues(calibStrings)

	fmt.Println(len(calibDigits))

	total := calibSum(calibDigits)

	fmt.Println(total)
}
