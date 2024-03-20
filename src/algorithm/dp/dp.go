package dp

// 数字三角形问题：动态规划（DP）求解
//   - https://www.luogu.com.cn/problem/P1216
func DigitalTrianglesProblemDP(nums []int) int {
	nums2d := make([][]int, 0, 10)
	for N, n := 0, 1; N < len(nums); n += 1 {
		N += n
		if N > len(nums) {
			panic("Nums Error")
		}
		temp := make([]int, n)
		copy(temp, nums[N-n:N])
		nums2d = append(nums2d, temp)
	}
	return digitalTrianglesProblemDP(nums2d)
}

func digitalTrianglesProblemDP(nums2d [][]int) int {
	dp := make([][]int, len(nums2d))
	for i := range dp {
		dp[i] = make([]int, len(nums2d[i]))
		copy(dp[i], nums2d[i])
	}
	for i := len(nums2d) - 1; i > 0; i-- {
		for j := 0; j < len(dp[i])-1; j++ {
			dp[i-1][j] = nums2d[i-1][j] + max(dp[i][j], dp[i][j+1])
		}
	}
	return dp[0][0]
}

// 采药问题：深度优先搜索（DFS）
//   - https://www.luogu.com.cn/problem/P1048
func GatherHerbsProblemDFS(arr2d [][]int) int {
	return gatherHerbsProblemDFS(1, 0, arr2d)
}

func gatherHerbsProblemDFS(i, t int, arr2d [][]int) int {
	if i > arr2d[0][1] {
		return 0
	}
	m1 := gatherHerbsProblemDFS(i+1, t, arr2d)
	m2 := 0
	if t := t + arr2d[i][0]; t <= arr2d[0][0] {
		m2 = gatherHerbsProblemDFS(i+1, t, arr2d) + arr2d[i][1]
	}
	return max(m1, m2)
}

// 采药问题：记忆化搜索（MS）
//   - https://www.luogu.com.cn/problem/P1048
func GatherHerbsProblemMS(arr2d [][]int) int {
	mem := make([][]*int, arr2d[0][1]+2)
	for i := range mem {
		mem[i] = make([]*int, arr2d[0][0]+1)
	}
	return gatherHerbsProblemMS(1, 0, mem, arr2d)
}

func gatherHerbsProblemMS(i, t int, mem [][]*int, arr2d [][]int) int {
	if mem[i][t] != nil {
		return *mem[i][t]
	}
	if i > arr2d[0][1] {
		return 0
	}
	m1 := gatherHerbsProblemMS(i+1, t, mem, arr2d)
	m2 := 0
	if t := t + arr2d[i][0]; t <= arr2d[0][0] {
		m2 = gatherHerbsProblemMS(i+1, t, mem, arr2d) + arr2d[i][1]
	}
	mem[i][t] = new(int)
	*mem[i][t] = max(m1, m2)
	return *mem[i][t]
}

// 采药问题：动态规划（DP）
//   - https://www.luogu.com.cn/problem/P1048
func GatherHerbsProblemDP(arr2d [][]int) int {
	dp := make([][]int, arr2d[0][1]+1)
	for i := range dp {
		dp[i] = make([]int, arr2d[0][0]+1)
	}
	for i := arr2d[0][1] - 1; i >= 0; i-- {
		for j := arr2d[0][0]; j >= 0; j-- {
			dp[i][j] = dp[i+1][j]
			if j <= arr2d[0][0]-arr2d[i][0] {
				dp[i][j] = max(dp[i][j], dp[i+1][j+arr2d[i][0]]+arr2d[i][1])
			}
		}
	}
	return dp[0][0]
}
