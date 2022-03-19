package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicate(t *testing.T) {
	duplicate := containsDuplicate([]int{1, 2, 3, 1})
	assert.Equal(t, false, duplicate)
}
