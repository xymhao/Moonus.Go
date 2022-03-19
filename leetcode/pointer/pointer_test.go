package pointer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	squares := sortedSquares([]int{-4, -1, 0, 3, 10})
	assert.Equal(t, 0, squares[0])
	assert.Equal(t, 1, squares[1])
	assert.Equal(t, 9, squares[2])
	assert.Equal(t, 16, squares[3])
	assert.Equal(t, 100, squares[4])
}

func Test2(t *testing.T) {
	squares := sortedSquares([]int{-5, -3, -2, -1})
	assert.Equal(t, 1, squares[0])
	assert.Equal(t, 4, squares[1])
	assert.Equal(t, 9, squares[2])
	assert.Equal(t, 25, squares[3])
}

func TestRotate2(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	rotate2(ints, 2)
	assert.Equal(t, 6, ints[0])
	assert.Equal(t, 7, ints[1])
	assert.Equal(t, 1, ints[2])
	assert.Equal(t, 2, ints[3])
	assert.Equal(t, 3, ints[4])

	ints2 := []int{1, 2, 3, 4, 5, 6}
	rotate2(ints2, 2)
	assert.Equal(t, 5, ints2[0])
	assert.Equal(t, 6, ints2[1])
	assert.Equal(t, 4, ints2[5])

}

func TestGcd(t *testing.T) {
	assert.Equal(t, 2, gcd(2, 6))
	assert.Equal(t, 1, gcd(2, 7))
}

func TestMoveZeroes(t *testing.T) {
	b := []int{0, 1, 0, 3, 5}
	moveZeroes(b)
	assert.Equal(t, 1, b[0])
	assert.Equal(t, 3, b[1])
	assert.Equal(t, 5, b[2])
	assert.Equal(t, 0, b[3])
}
