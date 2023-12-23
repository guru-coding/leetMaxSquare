package leetMaxSquare

func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([]int, n+1)
	mem1 := 0
	mem2 := 0
	maxsqlen := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			mem1 = dp[j]
			if matrix[i-1][j-1] == 0 {
				dp[j] = 0
				continue
			}
			dp[j] = min(min(dp[j-1], mem2), dp[j]) + 1
			maxsqlen = max(maxsqlen, dp[j])
			mem2 = mem1
		}
	}
	return maxsqlen * maxsqlen
}
