package _206_reverse_linked_list

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func ReverseList(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
