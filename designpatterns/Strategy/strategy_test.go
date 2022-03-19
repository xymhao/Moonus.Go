package Strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrategy(t *testing.T) {
	operator := Operator{
		Strategy: &add{},
	}
	result := operator.Calculator(1, 2)
	assert.Equal(t, 3, result)

	operator.Strategy = &reduce{}

	result = operator.Calculator(3, 1)
	assert.Equal(t, 2, result)
}
