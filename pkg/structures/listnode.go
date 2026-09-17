package structures

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	sb := strings.Builder{}
	for cur := l; cur != nil; cur = cur.Next {
		if sb.Len() > 0 {
			sb.WriteString(" -> ")
		}
		sb.WriteString(strconv.Itoa(cur.Val))
	}

	return sb.String()
}

func CreateListNode(nums []int) *ListNode {
	dummy := &ListNode{}
	op := dummy
	for i := 0; i < len(nums); i++ {
		op.Next = &ListNode{
			Val: nums[i],
		}
		op = op.Next
	}
	return dummy.Next
}
