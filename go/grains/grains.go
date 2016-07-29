package grains

import "errors"

var err = errors.New("Incorrect input")

func Square(input int) (uint64, error) {
	if input <= 0 || input > 64 {
		return 0, err
	}
	// moving by one bit gives the same result as multiplication by 2, just faster
	return 1 << uint(input-1), nil
}

func Total() uint64 {
	return 1<<64 - 1
}
