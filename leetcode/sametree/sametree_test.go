package sametree

import "testing"

func TestIsSameTree(t *testing.T) {
	node := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	tree := isSameTree(&node, &node)

	if !tree {
		t.Errorf("校验错误")
	}

	node2 := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2}}
	tree2 := isSameTree(&node, &node2)

	if tree2 {
		t.Errorf("校验错误")
	}
}
