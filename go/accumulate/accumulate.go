package accumulate

const testVersion = 1

func Accumulate(input []string, function func(string) string) []string {
	for i := range input {
		input[i] = function(input[i])
	}
	return input
}
