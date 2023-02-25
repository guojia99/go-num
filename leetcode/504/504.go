package _04

import "strconv"

// 给定一个整数 num，将其转化为 7 进制，并以字符串形式输出。
//
// 示例 1:
//
// 输入: num = 100
// 输出: "202"
// 示例 2:
//
// 输入: num = -7
// 输出: "-10"
//
// 提示：
//
// -107 <= num <= 107
func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(-num)
	}
	if num < 7 {
		return strconv.Itoa(num)
	}
	return convertToBase7(num/7) + convertToBase7(num%7)
}
