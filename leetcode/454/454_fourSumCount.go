package _54

// 给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，请你计算有多少个元组 (i, j, k, l) 能满足：
//
//0 <= i, j, k, l < n
//nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/4sum-ii
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func maxMin(a, b int, isMax bool) int {
	if isMax && a > b {
		return a
	}
	if !isMax && a < b {
		return a
	}
	return b
}
