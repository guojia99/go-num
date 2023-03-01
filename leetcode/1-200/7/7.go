/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/3/1 下午8:49.
 * Author:  guojia(https://github.com/guojia99)
 */

package _7

// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。
//
// 示例 1：
//
// 输入：x = 123
// 输出：321
// 示例 2：
//
// 输入：x = -123
// 输出：-321
// 示例 3：
//
// 输入：x = 120
// 输出：21
// 示例 4：
//
// 输入：x = 0
// 输出：0
//
// 提示：
//
// -2^31 <= x <= 2^31 - 1

func reverse2(x int) int {
	t := 0
	for x != 0 {
		t = t*10 + x%10
		x /= 10
	}
	if t > 1<<31-1 || t < -1<<31 {
		return 0
	}
	return t
}

func reverse(x int) int {
	isMinus := x < 0
	if isMinus {
		x = -x
	}

	var s []int
	for {
		in := x % 10
		if length := len(s); length != 0 || (length == 0 && in != 0) {
			s = append(s, in)
			if x < 10 {
				break
			}
		}

		x /= 10
		if x < 10 {
			s = append(s, x)
			break
		}
	}

	out := 0
	f := 1
	for i := len(s) - 1; i >= 0; i-- {
		out += f * s[i]
		f *= 10
	}
	if out > 2147483647 {
		return 0
	}

	if isMinus {
		out = -out
	}

	return out
}
