package grains

import "errors"

// Square number of grains on the square
func Square(grains int) (uint64, error) {
	if grains <= 0 || grains > 64 {
		return 0, errors.New("Error grains out of range")
	}

	return 1 << (grains - 1), nil
}

// Total number of grains on a chessboard
func Total() uint64 {
	return (1 << 64) - 1
}
