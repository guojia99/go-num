package _3

func lengthOfLongestSubstring(s string) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	out := 0
	start := 0
	cnt := [128]rune{}
	for idx, ss := range s {
		cnt[ss] += 1
		for cnt[ss] > 1 {
			cnt[s[start]] -= 1
			start += 1
		}
		out = max(out, idx-start+1)
	}
	return out
}
