package triangle

import "math"

const testVersion = 2

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

var (
	pinf = math.Inf(1)
	ninf = math.Inf(-1)
)

func badNumber(n float64) bool {
	if n <= 0 || math.IsNaN(n) || n == pinf || n == ninf {
		return true
	}
	return false
}

func KindFromSides(a, b, c float64) Kind {
	if badNumber(a) || badNumber(b) || badNumber(c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a+b < c || b+c < a || c+a < b {
		return NaT
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
