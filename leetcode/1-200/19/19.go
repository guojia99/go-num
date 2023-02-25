package _9

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
// 前后双指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummyHead = &ListNode{}
	dummyHead.Next = head
	var f, s = dummyHead, dummyHead
	// 这里将前指针指向n个之后
	for ; n != 0; n-- {
		f = f.Next
	}
	f = f.Next
	// 这里同步将前针和后针都往后推，直到前指针为空时，则后指针刚好为倒数第n个
	for f != nil {
		f, s = f.Next, s.Next
	}
	// 此时前后替换即可
	s.Next = s.Next.Next
	return dummyHead.Next
}
