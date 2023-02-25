package _39

// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length
//
// 链接：https://leetcode.cn/problems/sliding-window-maximum
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	var queue []int
	var out []int

	for r := 0; r < len(nums); r++ {

		// 队列不空且考察元素大于队尾元素 移除队尾元素
		for len(queue) > 0 {
			qBackIdx := len(queue) - 1
			if nums[r] > nums[queue[qBackIdx]] {
				queue = queue[:qBackIdx]
				continue
			}
			break
		}
		// 这时候考察元素一定小于队尾元素
		queue = append(queue, r)

		l := r - k + 1
		// queue 存放的是下标 left 是左极限 滑动要将数组弹出
		for queue[0] < l {
			queue = queue[1:]
		}

		// right = k-1 时候 说明第一个窗口形成 那么可以加入res
		// 之后right 每次++ out
		if r >= k-1 {
			out = append(out, nums[queue[0]])
		}
	}
	return out
}

// 暴力 超时
func maxSlidingWindowBF(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}

	var out []int
	var l, r = 0, k
	for ; r <= len(nums); l, r = l+1, r+1 {
		max := nums[l]
		for i := l; i < r; i++ {
			if nums[i] >= max {
				max = nums[i]
			}
		}
		out = append(out, max)
	}
	return out
}
