package chapter1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHanoi(t *testing.T) {
	assert := assert.New(t)
	tc := []string{"1->2", "1->3", "2->3", "1->2", "3->1", "3->2", "1->2"}
	assert.Equal(tc, Hanoi(3))
}
