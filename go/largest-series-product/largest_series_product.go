package lsproduct

import "errors"

const testVersion = 4

func LargestSeriesProduct(input string, span int) (int64, error) {
	maxProduct := int64(0)
	tmpProduct := int64(1)
	if span < 0 || span > len(input) {
		return -1, errors.New("Wrong span")
	}
	if span == 0 {
		return 1, nil
	}
	digits := make([]uint8, len(input))
	for i := range input {
		ch := input[i]
		if ch < '0' || ch > '9' {
			return 0, errors.New("Unknown char")
		}
		digits[i] = ch - '0'
	}
	usedDigits := 0
	for i := 0; i < len(input); i++ {
		if digits[i] == 0 {
			tmpProduct = 1
			usedDigits = 0
			continue
		}
		if usedDigits >= span && digits[i-span] != 0 {
			tmpProduct /= int64(digits[i-span])
		}
		tmpProduct *= int64(digits[i])
		usedDigits++
		if tmpProduct > maxProduct && usedDigits >= span {
			maxProduct = tmpProduct
		}
	}
	return maxProduct, nil
}
