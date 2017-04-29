package pythagorean

import "math"

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (ts []Triplet) {
	a, b, c2 := min, min, max
	var c float64
	for a = min; a <= max; a++ {
		for b = a; b <= max; b++ {
			c2 = a*a + b*b
			c = math.Sqrt(float64(c2))
			if int(c) > max {
				break
			}
			if math.Mod(c, 1) == 0.0 {
				ts = append(ts, Triplet{a, b, int(c)})
			}
		}
	}
	return
}

func Sum(p int) (ts []Triplet) {
	for a := 1; a < p; a++ {
		for b := a; b < p-a; b++ {
			c := p - a - b
			if c < b {
				break
			}
			if a*a+b*b == c*c {
				ts = append(ts, Triplet{a, b, c})
			}
		}
	}
	return
}
