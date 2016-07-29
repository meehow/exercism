package secret

func Handshake(code int) []string {
	operations := make([]string, 0, 4)
	if code <= 0 {
		return operations
	}
	if code&1 > 0 {
		operations = append(operations, "wink")
	}
	if code&2 > 0 {
		operations = append(operations, "double blink")
	}
	if code&4 > 0 {
		operations = append(operations, "close your eyes")
	}
	if code&8 > 0 {
		operations = append(operations, "jump")
	}
	if code&16 > 0 {
		for left, right := 0, len(operations)-1; left < right; {
			operations[left], operations[right] = operations[right], operations[left]
			left++
			right--
		}
	}
	return operations
}
