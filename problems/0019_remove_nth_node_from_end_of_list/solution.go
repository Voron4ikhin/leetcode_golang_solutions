package _019_remove_nth_node_from_end_of_list

import (
	"github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"
)

func RemoveNthFromEnd(head *structures.ListNode, n int) *structures.ListNode {
	dummy := &structures.ListNode{Val: 0}
	dummy.Next = head

	slow, fast := dummy, dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
