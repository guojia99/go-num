package _0

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
// 有效字符串需满足：
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 每个右括号都有一个对应的相同类型的左括号。
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
// 示例 2：
//
// 输入：s = "()[]{}"
// 输出：true
// 示例 3：
//
// 输入：s = "(]"
// 输出：false
func isValid(s string) bool {
	var lasts []rune

	for _, ss := range s {
		switch ss {
		case '[', '{', '(':
			lasts = append(lasts, ss)
		case ']', '}', ')':
			length := len(lasts)
			if length == 0 {
				return false
			}
			if !((lasts[length-1] == '(' && ss == ')') ||
				(lasts[length-1] == '{' && ss == '}') ||
				(lasts[length-1] == '[' && ss == ']')) {
				return false
			}
			lasts = lasts[:length-1]
		}
	}

	return len(lasts) == 0
}
