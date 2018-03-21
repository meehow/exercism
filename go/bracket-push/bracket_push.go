package brackets

var pair = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

// Bracket checks if bracket chain is valid
func Bracket(input string) (bool, error) {
	brackets := make([]rune, 0, len(input)/2)
	for _, c := range input {
		switch c {
		case '[', '{', '(':
			brackets = append(brackets, c)
		case ']', '}', ')':
			if len(brackets) == 0 {
				return false, nil
			}
			if pair[brackets[len(brackets)-1]] != c {
				return false, nil
			}
			brackets = brackets[:len(brackets)-1]
		}
	}
	return len(brackets) == 0, nil
}
