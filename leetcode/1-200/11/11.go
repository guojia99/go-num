package _1

// maxArea
// 给定一个长度为的整数数组 height。有n条垂线，第i条线的两个端点是（i，）和（i，height[i]）。
// 找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
// 返回容器可以储存的最大水量
// 链接：https://leetcode.cn/problems/container-with-most-water
// n == height.length
// 2 <= n <= 105
// 0 <= height[i] <= 104

// - 找到最大的那个面积
// - 取决于两个点的高度和距离

func maxArea(height []int) int {
	var l, r = 0, len(height) - 1
	var res = 0

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for l < r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l += 1
			continue
		}
		res = max(res, height[r]*(r-l))
		r -= 1
	}
	return res
}
