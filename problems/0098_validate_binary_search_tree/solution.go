package _098_validate_binary_search_tree

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func IsValidBST(root *structures.TreeNode) bool {
	return recValidate(root, nil, nil)
}

func recValidate(n, min, max *structures.TreeNode) bool {
	if n == nil {
		return true
	}

	if min != nil && min.Val >= n.Val {
		return false
	}
	if max != nil && max.Val <= n.Val {
		return false
	}

	return recValidate(n.Left, min, n) && recValidate(n.Right, n, max)
}
