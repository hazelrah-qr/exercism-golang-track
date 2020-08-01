package encode

import "strings"

// RunLengthDecode returns the encoded string
// Example: AAB => 2A1B
func RunLengthDecode(input string) string {
	data := strings.ToLower(input)
	peek(data[0], data[1:])
	return ""
}

func peek(c byte, future string) {
	if c == future[0] {

	}
}
