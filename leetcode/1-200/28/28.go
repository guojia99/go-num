package _8

// 给你两个字符串haystack 和 needle ，
// 请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。如果needle 不是 haystack 的一部分，则返回 -1 。
// 链接：https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string
// 1 <= haystack.length, needle.length <= 104
// haystack 和 needle 仅由小写英文字符组成
func strStrDP(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		var not bool
		for j := range needle {
			if haystack[i+j] != needle[j] {
				not = true
			}
		}
		if !not {
			return i
		}
	}
	return -1
}
