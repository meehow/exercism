package sublist

// Relation can be one of `sublist`, `superlist`, `equal` or `unequal`
type Relation string

func isUnequal(shorterList, longerList []int) bool {
	for i := 0; i <= len(longerList)-len(shorterList); i++ {
		equal := true
		for j := 0; j < len(shorterList); j++ {
			if shorterList[j] != longerList[i+j] {
				equal = false
				break
			}
		}
		if equal {
			return false
		}
	}
	return true
}

// Sublist compares two given lists
func Sublist(listOne, listTwo []int) Relation {
	if len(listOne) < len(listTwo) {
		if isUnequal(listOne, listTwo) {
			return "unequal"
		}
		return "sublist"
	}
	if len(listOne) > len(listTwo) {
		if isUnequal(listTwo, listOne) {
			return "unequal"
		}
		return "superlist"
	}
	for i := 0; i < len(listOne); i++ {
		if listOne[i] != listTwo[i] {
			return "unequal"
		}
	}
	return "equal"
}
