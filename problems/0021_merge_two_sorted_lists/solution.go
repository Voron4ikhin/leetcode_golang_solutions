package _021_merge_two_sorted_lists

import (
	"github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"
)

func MergeTwoLists(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	blank := &structures.ListNode{}
	op := blank

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}
		op = op.Next
	}

	if list1 != nil {
		op.Next = list1
	}

	if list2 != nil {
		op.Next = list2
	}

	return blank.Next
}
