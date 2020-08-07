package grains

import "errors"

// Square number of grains on the square
func Square(grains int) (uint64, error) {
	if grains <= 0 {
		return 0, errors.New("Error")
	}

	if grains > 64 {
		return 0, errors.New("Too large number")
	}
	result := uint64(1) << (grains - 1)
	return result, nil
}

// Total number of grains on a chessboard
func Total() uint64 {
	result, _ := Square(64)
	result = uint64(result)<<1 - 1
	return result
}
