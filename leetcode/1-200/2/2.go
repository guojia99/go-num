package _2

import `fmt`

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	s := fmt.Sprintf("%d", l.Val)
	if l.Next != nil {
		return s + l.Next.String()
	}
	return s
}

// addTwoNumbers
// 给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 链接：https://leetcode.cn/problems/add-two-numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head

	val := 0
	for l1 != nil || l2 != nil || val != 0 {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		res := (v1 + v2 + val) / 10
		val = (v1 + v2 + val) % 10
		cur.Val, val = val, res

		if l1 != nil || l2 != nil || val != 0 {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}
	return head
}
