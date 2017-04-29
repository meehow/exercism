package acronym

import "unicode"

const testVersion = 2

func Abbreviate(input string) string {
	previousIsLower := false
	previousIsLetter := false
	acronym := make([]rune, 0)
	for _, r := range input {
		if !previousIsLetter && unicode.IsLetter(r) {
			acronym = append(acronym, unicode.ToUpper(r))
		} else if previousIsLower && unicode.IsUpper(r) {
			acronym = append(acronym, r)
		}
		previousIsLower = unicode.IsLower(r)
		previousIsLetter = unicode.IsLetter(r)

	}
	return string(acronym)
}
