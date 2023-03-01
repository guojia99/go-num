/*
 * Copyright (c) 2023 guojia99 All rights reserved.
 * Created: 2023/2/25 上午11:35.
 * Author:  guojia(https://github.com/guojia99)
 */

package _5

// longestPalindrome
// 给你一个字符串 s，找到 s 中最长的回文子串。
//
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
//
// 示例 1：
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
//
// 输入：s = "cbbd"
// 输出："bb"
//
// 提示：
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
func longestPalindrome(s string) string {
	var start, end int

	for i := 0; i < len(s); {
		l, r := i, i

		for r < len(s)-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1

		for l > 0 && r < len(s)-1 && s[l-1] == s[r+1] {
			l -= 1
			r += 1
		}
		if end-start < r-l {
			start, end = l, r
		}
	}

	return s[start : end+1]
}
