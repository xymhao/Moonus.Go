package symmetrictree

import "testing"

func TestIsSymmetric(t *testing.T) {
	symmetric := isSymmetric(&TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}})

	if !symmetric {
		t.Errorf("失败")
	}
}
