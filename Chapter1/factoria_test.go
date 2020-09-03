package chapter1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorial1(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5*4*3*2*1, facItr1(5))
}

func TestFactorial2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(5*4*3*2*1, facItr2(5))
}
