/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/25 上午11:44.
 * Author:  guojia(https://github.com/guojia99)
 */

package _4

import "sort"

//findMedianSortedArrays
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
//算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
//示例 1：
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//示例 2：
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//提示：
//
//nums1.length == m
//nums2.length == n
//0 <= m <= 1000
//0 <= n <= 1000
//1 <= m + n <= 2000
//-106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	l := len(nums1)
	if l == 0 {
		return 0
	} else if l == 1 {
		return float64(nums1[0])
	}

	if l%2 == 0 {
		return float64(nums1[l/2])/2 + float64(nums1[l/2-1])/2
	}
	return float64(nums1[l/2])
}
