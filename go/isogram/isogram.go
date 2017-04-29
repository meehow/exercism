package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	chars := strings.ToLower(word)[:]
	for i, r1 := range chars {
		if unicode.IsLetter(r1) == false {
			continue
		}
		for _, r2 := range chars[i+1:] {
			if r1 == r2 {
				return false
			}
		}
	}
	return true
}
