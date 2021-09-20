package maxdepthoftree

import "testing"

func TestMax(t *testing.T) {
	node := TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 1}}}

	depth := maxDepth(&node)
	if depth != 3 {
		t.Errorf("error")
	}

	depth2 := Breadth(&node)
	if depth2 != 3 {
		t.Errorf("error")
	}
}
