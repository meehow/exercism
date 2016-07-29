package queenattack

import "errors"

var err = errors.New("err")

func CanQueenAttack(w, b string) (bool, error) {
	if w == b || b[1] > '8' {
		return false, err
	}
	if w[0] == b[0] || w[1] == b[1] || w[0]-b[0] == w[1]-b[1] || w[0]-b[0] == b[1]-w[1] {
		return true, nil
	}
	return false, nil
}
