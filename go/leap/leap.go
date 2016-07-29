package leap

const testVersion = 2

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		return !(year%100 == 0 && year%400 != 0)
	}
	return false
}
