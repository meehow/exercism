package brackets

var pair = map[rune]rune{
	'[': ']',
	'{': '}',
	'(': ')',
}

// Cell represents one cell of bracket chain
type Cell struct {
	Parent  *Cell
	Bracket rune
}

// Bracket checks if bracket chain is valid
func Bracket(input string) (bool, error) {
	parent := new(Cell)
	for _, c := range input {
		switch c {
		case '[', '{', '(':
			parent = &Cell{
				Parent:  parent,
				Bracket: c,
			}
		case ']', '}', ')':
			if parent.Parent == nil {
				return false, nil
			}
			if pair[parent.Bracket] != c {
				return false, nil
			}
			parent = parent.Parent
		}
	}
	return parent.Parent == nil, nil
}
