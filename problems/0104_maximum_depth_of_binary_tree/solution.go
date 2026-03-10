package _104_maximum_depth_of_binary_tree

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func MaxDepth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	return 1 + max(left, right)
}
