package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return initBst(nums, 0, len(nums)-1)
}

func initBst(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := TreeNode{Val: nums[mid]}

	root.Left = initBst(nums, left, mid-1)
	root.Right = initBst(nums, mid+1, right)
	return &root
}
