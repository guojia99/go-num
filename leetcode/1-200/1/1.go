package _1

import (
	`fmt`
	`sort`
)

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//
//你可以按任意顺序返回答案。
//
// 
//
//示例 1：
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//示例 2：
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//示例 3：
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
// 
//
//提示：
//
//2 <= nums.length <= 104
//-109 <= nums[i] <= 109
//-109 <= target <= 109
//只会存在一个有效答案
func twoSum(nums []int, target int) []int {
	var dp = make([][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = [2]int{i, nums[i]}
	}
	sort.Slice(dp, func(i, j int) bool {
		return dp[i][1] < dp[j][1]
	})
	fmt.Println(dp)

	var l, r = 0, len(nums) - 1
	for l < r {
		sum := dp[l][1] + dp[r][1]
		if sum == target {
			break
		} else if sum > target {
			r--
		} else {
			l++
		}
	}
	return []int{dp[l][0], dp[r][0]}
}
