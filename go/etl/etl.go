package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int, 40)
	for val, letters := range input {
		for _, letter := range letters {
			// it would be faster to just add 32 to a char code
			output[strings.ToLower(letter)] = val
		}
	}
	return output
}
