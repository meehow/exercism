package palindrome

const testVersion = 1

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	return pmin, pmax, nil
}
