package _2

// 给定n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
// 示例 1：
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
// 输入：height = [4,2,0,3,2,5]
// 输出：9
//
//	|
//
// | × × × × |
// | × × | × |
// | | × | | |
// | | × | | |
// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105
// 链接：https://leetcode.cn/problems/trapping-rain-water

func trap(height []int) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	// 两个指针l指向第一列，r指向最后一列，分别用于用变量维护当前指针的前缀和后缀最大值。
	// 若当前前缀最大值小于后缀最大值，那么当前列l就可以计算出积水高度，否则r列可以计算出积水高度。
	length := len(height)
	if length <= 2 {
		return 0
	}

	var out = 0
	var lMax, rMax = height[0], height[length-1]
	var l, r = 0, length - 1

	for l <= r {
		lH := height[l]
		rH := height[r]

		lMax = max(lMax, lH)
		rMax = max(rMax, rH)
		if lMax <= rMax {
			out += lMax - lH
			l++
		} else {
			out += rMax - rH
			r--
		}
	}
	return out
}
