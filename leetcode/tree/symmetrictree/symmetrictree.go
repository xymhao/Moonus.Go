package symmetrictree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	left := root.Left
	right := root.Right
	return check(left, right)
}

func check(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}

	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
