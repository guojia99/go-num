package _09

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
//
// 示例 1：
//
// 输入：target = 7, nums = [2,3,1,2,4,3]
// 输出：2
// 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
// 示例 2：
//
// 输入：target = 4, nums = [1,4,4]
// 输出：1
// 示例 3：
//
// 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
// 输出：0
//
// 提示：
//
// 1 <= target <= 109
// 1 <= nums.length <= 105
// 1 <= nums[i] <= 105
func minSubArrayLen(target int, nums []int) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	l, minLen := 0, len(nums)
	sum := 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for l <= r && sum > target {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
	}
	if l == 0 {
		return 0
	}
	return minLen
}
