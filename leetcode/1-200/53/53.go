package _3

/**
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
 子数组 是数组中的一个连续部分。
 示例 1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1] 的和最大，为6 。
 示例 2：
输入：nums = [1]
输出：1
 示例 3：
输入：nums = [5,4,-1,7,8]
输出：23
 提示：
 1 <= nums.length <= 10⁵
 -10⁴ <= nums[i] <= 10⁴
 进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。
 Related Topics 数组 分治 动态规划 👍 5682 👎 0
*/

// leetcode submit region begin(Prohibit modification and deletion)
func maxSubArray(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ans := 0
	out := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = max(nums[i], ans+nums[i])
		out = max(out, ans)
	}
	return out
}

// 执行耗时:92 ms,击败了20.29% 的Go用户
// 内存消耗:8.7 MB,击败了75.09% 的Go用户
func maxSubArrayMy(nums []int) int {
	if len(nums) <= 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	maxNum := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		maxNum = max(dp[i], maxNum)
	}
	return maxNum
}

//leetcode submit region end(Prohibit modification and deletion)
