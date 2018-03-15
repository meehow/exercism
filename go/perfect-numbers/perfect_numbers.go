package perfect

import (
	"errors"
	"math"
)

type Classification uint8

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("ErrOnlyPositive")

func Classify(number int64) (Classification, error) {
	if number <= 0 {
		return ClassificationDeficient, ErrOnlyPositive
	}
	sum := AliquotSum(number)
	if sum == number {
		return ClassificationPerfect, nil
	} else if sum > number {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}

// func AliquotSumSlow(number int64) (sum int64) {
// 	for i := int64(1); i <= number/2; i++ {
// 		if number%i == 0 {
// 			sum += i
// 		}
// 	}
// 	return
// }

func AliquotSum(number int64) int64 {
	sum := int64(1)
	p := int64(math.Sqrt(float64(number)))
	k := int64(2)
	for number > 1 && k <= p {
		e := float64(1)
		for number%k == 0 {
			number /= k
			e++
		}
		if e > 1 {
			sum *= int64(math.Pow(float64(k), e)-1) / (k - 1)
		}
		k++
	}
	if number > 1 {
		sum *= (number*number - 1) / (number - 1)
	}
	return sum / 2
}
