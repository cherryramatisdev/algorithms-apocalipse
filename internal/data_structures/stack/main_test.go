package stack_test

import (
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/internal/data_structures/stack"
	"gotest.tools/assert"
)

func TestPush(t *testing.T) {
	list := stack.NewStack[int]()

	list.Push(5)
	list.Push(7)
	list.Push(9)

	assert.Equal(t, list.Length, 3)
}

func TestPop(t *testing.T) {
	list := stack.NewStack[int]()

	list.Push(5)
	list.Push(7)
	list.Push(9)

	assert.Equal(t, list.Length, 3)
	assert.Equal(t, list.Pop(), 9)
	assert.Equal(t, list.Length, 2)
}

func TestPeek(t *testing.T) {
	list := stack.NewStack[int]()

	list.Push(5)
	list.Push(7)
	list.Push(9)

	assert.Equal(t, list.Length, 3)
	assert.Equal(t, list.Peek(), 9)
	assert.Equal(t, list.Length, 3)
}
