package _049_group_anagrams

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func GetIntersectionNode(headA, headB *structures.ListNode) *structures.ListNode {
	lenA := GetNodeLen(headA)
	lenB := GetNodeLen(headB)
	a := Shift(headA, max(lenA-lenB, 0))
	b := Shift(headB, max(lenB-lenA, 0))
	for {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
}

func GetNodeLen(a *structures.ListNode) int {
	length := 0
	for a != nil {
		length++
		a = a.Next
	}
	return length
}

func Shift(a *structures.ListNode, steps int) *structures.ListNode {
	for steps > 0 {
		a = a.Next
		steps--
	}

	return a
}
