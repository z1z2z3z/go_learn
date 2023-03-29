/*
 * @Description: 请填写简介
 */
 package main

func isMatch(s string, p string) bool {

	//cccf
	// c*b*f
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (s[i-1] == p[j-2] || p[j-2] == '.') && dp[i-1][j])
			} else {
				if i >= 1 && j >= 1 {
					dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
				} else {
					if i == 0 && j != 0 {
						dp[0][j] = false
					} else if i != 0 && j == 0 {
						dp[i][0] = false
					} else {
						dp[0][0] = true
					}
				}
			}
		}
	}
	return dp[m][n]
}