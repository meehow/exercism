package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return "Fine. Be that way!"
	}
	if isShouting(s) {
		return "Whoa, chill out!"
	}
	if s[len(s)-1] == '?' {
		return "Sure."
	}
	return "Whatever."
}

func isShouting(s string) bool {
	shouting := false
	for _, r := range s {
		if !shouting && unicode.IsUpper(r) {
			shouting = true
		} else if unicode.IsLower(r) {
			return false
		}
	}
	return shouting
}
