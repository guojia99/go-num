package _8

import (
	"sort"
)

// fourSum
// 给你一个由 n 个整数组成的数组nums ，和一个目标值 target
// 找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]（若两个四元组元素一一对应，则认为两个四元组重复）：
// 0 <= a, b, c, d< n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案 。
// 链接：https://leetcode.cn/problems/4sum
func fourSum(nums []int, target int) [][]int {
	if len(nums) <= 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	res := make([][]int, 0)
	var temp = make(map[[4]int]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			l, r := j+1, len(nums)-1
			for l < r {
				count := nums[i] + nums[j] + nums[l] + nums[r]
				if count == target {
					m := [4]int{nums[i], nums[j], nums[l], nums[r]}
					if _, ok := temp[m]; !ok {
						temp[m] = struct{}{}
						res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					}
				}
				if count > target {
					r--
				} else {
					l++
				}
			}
		}
	}
	return res
}
