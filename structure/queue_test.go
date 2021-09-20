package structure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(1)

	dequeue := queue.Dequeue()
	assert.Equal(t, 1, dequeue)

}
