package luhn

import (
	"regexp"
	"strconv"
)

func Valid(input string) bool {
	reRemoveSpaces, _ := regexp.Compile(" ")
	input = reRemoveSpaces.ReplaceAllString(input, "")

	reFindNotDigits, _ := regexp.Compile("[^0-9]")
	if reFindNotDigits.MatchString(input) {
		return false
	}

	if input == "0" {
		return false
	}

	var acc, b int
	odd := len(input) % 2
	for i, rune := range input {
		char := string(rune)
		digit, _ := strconv.Atoi(char)
		// digit, _ := strconv.ParseInt(char, 10, 64)
		if i%2 == odd {
			b = 2 * digit
			if b > 9 {
				b -= 9
			}
			acc += b
		} else {
			acc += digit
		}
	}

	return acc%10 == 0
}
