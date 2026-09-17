package structures

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	sb := strings.Builder{}
	for l != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", l.Val))
		l = l.Next
	}

	return sb.String()
}
