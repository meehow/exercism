package slice

func All(n int, s string) []string {
	count := len(s) - n + 1
	series := make([]string, count)
	for i := 0; i < count; i++ {
		series[i] = s[i : i+n]
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
