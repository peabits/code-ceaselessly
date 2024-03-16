package dp

// 最长（严格）递增子序列问题：动态规划（DP）求解
//
//   - https://leetcode.cn/problems/longest-increasing-subsequence/description/
func LengthOfLISProblemDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ { // 状态转移
			if nums[j] < nums[i] { // 非严格为 <=
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	return maxLen
}

// 最长（严格）递增子序列问题：贪心算法（GA）求解
//
//   - https://leetcode.cn/problems/longest-increasing-subsequence/description/
func LengthOfLISProblemGA(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	seq := make([]int, 1, len(nums))
	seq[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > seq[len(seq)-1] { // 非严格为 >=
			seq = append(seq, nums[i])
		} else { // 查找并替换第一个比其大的数
			j := inorderSearch(seq, nums[i]) // 顺序搜索
			_ = binarySearch(seq, nums[i])   // 或者使用二分搜索插入
			seq[j] = nums[i]
		}
	}
	return len(seq)
}

func inorderSearch(arr []int, v int) (pos int) {
	for i, n := range arr {
		if n >= v {
			pos = i
			break
		}
	}
	return
}

func binarySearch(arr []int, v int) (pos int) {
	l, r := 0, len(arr)
	for l <= r {
		m := (l + r) >> 1
		if arr[m] < v {
			l = m + 1
			pos = l
		} else {
			r = m - 1
		}
	}
	return
}
