package cryptosquare

import (
	"math"
	"strings"
)

const testVersion = 2

func Normalize(input string) string {
	input = strings.ToLower(input)
	output := make([]byte, 0, len(input))
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			output = append(output, ch)
		}
	}
	return string(output)
}

func Encode(plainText string) string {
	normal := Normalize(plainText)
	edge := math.Sqrt(float64(len(normal)))
	width := int(math.Ceil(edge))
	height := int(edge)
	length := len(normal)
	output := make([]byte, 0, length)
	for i := 0; i < width; i++ {
		if i != 0 {
			output = append(output, ' ')
		}
		for j := 0; j <= height; j++ {
			idx := j*width + i
			if idx >= length {
				break
			}
			output = append(output, normal[idx])
		}
	}
	return string(output)
}
