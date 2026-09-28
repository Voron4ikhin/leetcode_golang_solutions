package _102_binary_tree_level_order_traversal

import "github.com/voron4ikhin/leetcode_golang_solutions/pkg/structures"

func levelOrder(root *structures.TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*structures.TreeNode{root}
	for len(queue) != 0 {
		preRes := make([]int, 0)
		newQueue := make([]*structures.TreeNode, 0)
		for _, v := range queue {
			preRes = append(preRes, v.Val)
			if v.Left != nil {
				newQueue = append(newQueue, v.Left)
			}
			if v.Right != nil {
				newQueue = append(newQueue, v.Right)
			}
		}
		result = append(result, preRes)
		queue = newQueue
	}
	return result
}
