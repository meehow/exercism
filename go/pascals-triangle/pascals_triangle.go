package pascal

func Triangle(n int) [][]int {
	t := make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, i+1)
		t[i][0] = 1
		t[i][i] = 1
		for j := 1; j < i; j++ {
			t[i][j] = t[i-1][j-1] + t[i-1][j]
		}
	}
	return t
}
