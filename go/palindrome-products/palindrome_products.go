package palindrome

import (
	"errors"
)

const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func init() {

}

func isPalindrome(number int) bool {
	n := number
	tmp := 0
	for n != 0 {
		tmp = tmp*10 + n%10
		n /= 10
	}
	return tmp == number
}

// Products finds palindrome products
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}
	for i := fmax * 2; i >= fmin; i-- {
		x := i/2 + i%2
		y := x - i%2
		for {
			if x > fmax || y < fmin {
				break
			}
			prod := x * y
			if (prod <= pmin.Product || prod >= pmax.Product) && isPalindrome(prod) {
				if pmax.Product == 0 {
					pmin.Product = prod
					pmin.Factorizations = [][2]int{{y, x}}
					pmax.Product = prod
					pmax.Factorizations = [][2]int{{y, x}}
				} else if prod < pmin.Product {
					pmin.Product = prod
					pmin.Factorizations = [][2]int{{y, x}}
				} else if prod == pmin.Product {
					pmin.Factorizations = append(pmin.Factorizations, [2]int{y, x})
				} else if prod > pmax.Product {
					pmax.Product = prod
					pmax.Factorizations = [][2]int{{y, x}}
				} else if prod == pmax.Product {
					pmax.Factorizations = append(pmax.Factorizations, [2]int{y, x})
				}
			}
			y--
			x++
		}
	}
	if pmin.Product == 0 && pmax.Product == 0 {
		err = errors.New("no palindromes...")
	}
	return
}
