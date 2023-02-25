package _2

import (
	`reflect`
	`testing`
)

func GetListNodeByList(in []int) *ListNode {
	var (
		node *ListNode
		next *ListNode
	)
	for _, val := range in {
		newNode := &ListNode{Val: val}
		if node == nil {
			node = newNode
			next = node
		} else {
			next.Next = newNode
			next = node
		}
	}
	return node
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "1",
			args: args{
				l1: GetListNodeByList([]int{2, 4, 3}),
				l2: GetListNodeByList([]int{5, 6, 4}),
			},
			want: GetListNodeByList([]int{7, 0, 8}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
