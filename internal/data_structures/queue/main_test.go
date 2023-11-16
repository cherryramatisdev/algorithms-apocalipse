package queue_test

import (
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/internal/data_structures/queue"
	"gotest.tools/assert"
)

func TestEnqueue(t *testing.T) {
	list := queue.NewQueue[int]()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assert.Equal(t, list.Length, 3)
}

func TestDeque(t *testing.T) {
	list := queue.NewQueue[int]()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assert.Equal(t, list.Length, 3)
	assert.Equal(t, list.Deque(), 5)
	assert.Equal(t, list.Length, 2)
}

func TestPeek(t *testing.T) {
	list := queue.NewQueue[int]()

	list.Enqueue(11)

	assert.Equal(t, list.Length, 1)
	assert.Equal(t, list.Peek(), 11)
}
