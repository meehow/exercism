package hamming

import "errors"

const testVersion = 4

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Not equal length")
	}
	dist := 0
	// using indexes is 2 times faster than using range
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
