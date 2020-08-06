package luhn

import (
	"strings"
	"unicode"
)

// Valid determines wether or not a provided string is a valid Luhn number
func Valid(input string) bool {
	trimmedString := strings.ReplaceAll(input, " ", "")
	if len(trimmedString) <= 1 {
		return false
	}

	multiply := len(trimmedString) % 2
	sum := 0
	for index, x := range trimmedString {

		if !unicode.IsDigit(x) {
			return false
		}

		digit := int(x - '0')
		if index%2 == multiply {
			digit = 2 * digit
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}
