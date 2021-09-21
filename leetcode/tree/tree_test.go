package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTTree(t *testing.T) {
	bst := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	if bst.Val != 0 {
		t.Errorf("error")
	}

	assert.Equal(t, -10, bst.Left.Val)
	assert.Equal(t, -3, bst.Left.Right.Val)
	assert.Equal(t, 5, bst.Right.Val)
	assert.Equal(t, 9, bst.Right.Right.Val)
}
