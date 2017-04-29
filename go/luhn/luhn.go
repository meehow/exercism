package luhn

const testVersion = 2

func Valid(input string) bool {
	sum := 0
	length := 0
	for i := len(input) - 1; i >= 0; i-- {
		ch := input[i]
		if ch >= '0' && ch <= '9' {
			num := int(ch - '0')
			if length%2 == 1 {
				num *= 2
				if num > 9 {
					num -= 9
				}
			}
			sum += num
			length++
		} else if ch == ' ' {
			continue
		} else {
			return false
		}
	}
	if length < 2 {
		return false
	}
	return sum%10 == 0
}
