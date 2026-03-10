package _141_linked_list_cycle

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func HasCycle(head *structures.ListNode) bool {
	slow, fast := head, head
	if fast == nil {
		return false
	}
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
