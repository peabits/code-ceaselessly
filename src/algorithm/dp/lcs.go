package dp

// 最长公共子序列问题：动态规划（DP）求解
//
//   - https://leetcode.cn/problems/longest-common-subsequence/description/
func LongestCommonSubsequenceProblemDP(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- { // 状态转移
			if s1[i] == s2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1]) // 即：max(dp[i+1][j], dp[i][j+1], dp[i+1][j+1]), 其中 dp[i+1][j+1] 最小，故省去
			}
		}
	}
	return dp[0][0]
}

// 优化方式：滚动数组
//
// :TODO
