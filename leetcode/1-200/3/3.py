class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        max_len, start = 0, 0
        mask = 0
        i = -1
        for st in s:
            i += 1
            ss = ord(st)
            if (mask & ss) != 1 << (ss + 97) and i != len(s) - 1:
                mask = mask | 1 << (ss + 97)
                continue
            if i == len(s) - 1:
                max_len = max(i - start + 1, max_len)
            else:
                max_len = max(i - start, max_len)
            start, mask = i + 1, 0

        return max_len


print(Solution().lengthOfLongestSubstring("abcabcbb"))
print(Solution().lengthOfLongestSubstring(" "))