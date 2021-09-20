package maxdepthoftree

import "Moonus.Go/structure"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

// Breadth Breadth-first traversal
// 广度优先遍历
func Breadth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue structure.Queue
	queue.Enqueue(root)
	ans := 0
	for queue.Len > 0 {
		var size = queue.Len
		for size > 0 {
			node := queue.Dequeue().(*TreeNode)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
			size--
		}
		ans++
	}
	return ans
}
