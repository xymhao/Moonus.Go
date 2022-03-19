package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTowSum(t *testing.T) {
	sum := twoSum([]int{3, 2, 4}, 6)
	assert.Equal(t, 1, sum[0])
	assert.Equal(t, 2, sum[1])
}

func TestMerge(t *testing.T) {
	//merge([]int{1,2,3,0,0,0}, 3, []int{2,5,6}, 3)
	merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3)
}
