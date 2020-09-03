package chapter1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonozi(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(0, Fibonazi(1))
	assert.Equal(1, Fibonazi(2))
	assert.Equal(1, Fibonazi(3))
	assert.Equal(2, Fibonazi(4))
	assert.Equal(3, Fibonazi(5))
	assert.Equal(5, Fibonazi(6))
	assert.Equal(8, Fibonazi(7))
	assert.Equal(13, Fibonazi(8))
}
