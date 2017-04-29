package sieve

const testVersion = 1

func Sieve(limit int) []int {
	primes := []int{2}
sieve:
	for i := 3; i < limit; i += 2 {
		for _, prime := range primes {
			if i%prime == 0 {
				continue sieve
			}
		}
		primes = append(primes, i)
	}
	return primes
}
