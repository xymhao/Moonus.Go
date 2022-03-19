package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	i := search([]int{-1, 0, 3, 5, 9, 12}, 10)

	assert.Equal(t, -1, i)
}

func TestName(t *testing.T) {
	firstBadVersion(1)

}
