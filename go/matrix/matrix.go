package matrix

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// Matrix stores matrix data
type Matrix [][]int

// New builds new matrix from input data
func New(input string) (*Matrix, error) {
	if input[len(input)-1] == '\n' {
		return nil, errors.New("Last row empty")
	}
	m := make(Matrix, 0)
	size := 0
	var err error
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		numbers := strings.Split(text, " ")
		if size == 0 {
			size = len(numbers)
		} else if len(numbers) != size {
			return nil, errors.New("Wrong row size")
		}
		row := make([]int, size)
		for i, n := range numbers {
			row[i], err = strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
		}
		m = append(m, row)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Rows returns rows of the matrix
func (m Matrix) Rows() [][]int {
	m2 := make([][]int, len(m))
	for i := range m {
		m2[i] = make([]int, len(m[i]))
		copy(m2[i], m[i])
	}
	return m2
}

// Cols returns columns of the matric
func (m Matrix) Cols() [][]int {
	m2 := make([][]int, len(m[0]))
	for i := range m2 {
		m2[i] = make([]int, len(m))
		for j := range m2[i] {
			m2[i][j] = m[j][i]
		}
	}
	return m2
}

// Set sets value on the matrix at given coordinates
func (m Matrix) Set(r, c, val int) bool {
	if r >= len(m) || c >= len(m[0]) || r < 0 || c < 0 {
		return false
	}
	m[r][c] = val
	return true
}
