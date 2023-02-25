package _6

import "sort"

// threeSumClosest
// 给你一个长度为 n 的整数数组nums和 一个目标值target。请你从 nums 中选出三个整数，使它们的和与target最接近。
// 返回这三个数的和。
// 假定每组输入只存在恰好一个解。
// 链接：https://leetcode.cn/problems/3sum-closest
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
//
//	 先将数组从小到大排序，便于微调 sum 的大小。
//		从左到右遍历，先固定一个数，剩下的部分，用头尾双指针扫描
//		如果 sum 大于目标值，就右指针左移，使 sum 变小，否则左指针右移，sum 变大。
//		再看 abs(sum - target) 是否比之前更小了，如果是，将当前 sum 更新给 res
//		遍历结束，就有了最接近目标值的 sum
func threeSumClosest(nums []int, target int) int {
	intAbs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			count := nums[i] + nums[l] + nums[r]
			if intAbs(target-count) < intAbs(target-res) {
				res = count
			}
			if count < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
