package summultiples

func MultipleSummerSlow(divisors ...int) func(int) int {
	return func(input int) int {
		multiples := 0
		for i := 1; i < input; i++ {
			for _, divisor := range divisors {
				if i%divisor == 0 {
					multiples += i
					break
				}
			}
		}
		return multiples
	}
}

func MultipleSummer(divisors ...int) func(int) int {
	return func(input int) int {
		multiples := 0
		combinations := make([][]int, 0)
		for i := 0; i < len(divisors); i++ {
			for j := 0; j < len(divisors); j++ {
				if i != j && divisors[j] != 0 && divisors[i]%divisors[j] == 0 {
					divisors[i] = 0
				}
			}
		}
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			multiples += Sum(input, divisor)
			for i := len(combinations) - 1; i >= 0; i-- {
				combinations = append(combinations, append(combinations[i], divisor))
			}
			combinations = append(combinations, []int{divisor})
		}
		for _, combination := range combinations {
			if len(combination) > 1 {
				multiples -= Sum(input, Multi(combination))
			}
		}
		return multiples
	}
}

func Sum(input, divisor int) int {
	n := input / divisor
	if input%divisor == 0 {
		n--
	}
	return n * (n + 1) / 2 * divisor
}

func Multi(nums []int) int {
	result := 1
	for _, num := range nums {
		result *= num
	}
	return result
}
