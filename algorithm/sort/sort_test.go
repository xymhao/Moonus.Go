package sort

import "testing"
import "github.com/stretchr/testify/assert"

func TestBubbleSor(t *testing.T) {
	arr := []int{4, 5, 6, 3, 2, 1}
	bubbleSort(arr)
	for i := 0; i < 6; i++ {
		assert.Equal(t, i+1, arr[i])
	}
}
