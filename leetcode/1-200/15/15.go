package _5

import (
	"sort"
)

// threeSum
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
// 你返回所有和为 0 且不重复的三元组。
// 链接：https://leetcode.cn/problems/3sum
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	out := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 双指针开始
		l := i + 1
		r := len(nums) - 1
		for l < r {
			count := nums[i] + nums[l] + nums[r]
			if count == 0 {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				// 处理相同值
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if count < 0 {
				l++ // 右移
			} else {
				r-- // 左移
			}
		}
	}
	return out
}
