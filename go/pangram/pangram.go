package pangram

const testVersion = 1

func IsPangram(input string) bool {
	var letters uint32
	var letter uint8
	for i := 0; i < len(input); i++ {
		letter = input[i]
		if letter < 'A' || letter > 'z' {
			continue
		}
		// kind of ToLower()
		if letter < 'a' {
			letter += 'a' - 'A'
		}
		// double check in case we got something between 'Z' and 'a'
		if letter > 'z' {
			continue
		}
		letters |= 1 << (letter - 'a')
	}
	return letters == 67108863
}

// func IsPangram(input string) bool {
// 	letters := make(map[rune]bool)
// 	for _, l := range strings.ToLower(input) {
// 		if l < 'a' || l > 'z' {
// 			continue
// 		}
// 		letters[l] = true
// 	}
// 	return len(letters) == 26
// }
