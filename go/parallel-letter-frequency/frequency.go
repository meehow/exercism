package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(input []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range input {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}
	m := FreqMap{}
	for range input {
		for r, count := range <-c {
			m[r] += count
		}
	}
	return m
}
