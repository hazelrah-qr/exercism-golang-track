package luhn

import (
	"strconv"
	"strings"
)

// Valid determines wether or not a provided string is a valid Luhn number
func Valid(input string) bool {
	trimmedString := strings.ReplaceAll(input, " ", "")
	if len(trimmedString) <= 1 {
		return false
	}

	numbers := []int{}
	for _, x := range trimmedString {
		n, err := strconv.Atoi(string(x))

		if err != nil {
			return false
		}
		numbers = append(numbers, n)
	}

	for i := len(numbers) - 2; i >= 0; i -= 2 {
		newValue := 2 * numbers[i]
		if newValue > 9 {
			newValue = newValue - 9
		}
		numbers[i] = newValue
	}

	sum := 0
	for _, x := range numbers {
		sum += x
	}

	return sum%10 == 0
}
