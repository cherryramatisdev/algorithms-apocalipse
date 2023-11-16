package linked_lists_test

import (
	"testing"

	"github.com/cherryramatisdev/algorithms-apocalipse/internal/data_structures/linked_lists"
	"gotest.tools/assert"
)

func TestLinkedListIsEmpty(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	assert.Equal(t, linked_list.Length(), 0)
}

func TestLinkedListInsertAt(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	linked_list.InsertAt(1, 0)

	assert.Equal(t, linked_list.Length(), 1)

	linked_list.InsertAt(2, 1)

	assert.Equal(t, linked_list.Length(), 2)

	assert.Error(t, linked_list.InsertAt(2, 3), "Index out of bounds")
}

func TestLinkedListRemove(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	linked_list.Append(1)
	linked_list.Append(2)
	linked_list.Append(3)

	got, _ := linked_list.Remove(3)
	assert.Equal(t, got, 3)
	assert.Equal(t, linked_list.Length(), 2)
}

func TestLinkedListRemoveAt(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	linked_list.Append(1)
	linked_list.Append(2)
	linked_list.Append(3)

	got, _ := linked_list.RemoveAt(2)
	assert.Equal(t, got, 3)
	assert.Equal(t, linked_list.Length(), 2)
}

func TestLinkedListAppend(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	linked_list.Append(1)
	linked_list.Append(2)
	linked_list.Append(3)

	assert.Equal(t, linked_list.Length(), 3)
	assert.Equal(t, linked_list.Head.Next.Next.Data, 3)
}

func TestLinkedListPrepend(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	linked_list.Prepend(1)
	linked_list.Prepend(2)
	linked_list.Prepend(3)

	assert.Equal(t, linked_list.Length(), 3)
	assert.Equal(t, linked_list.Head.Data, 3)
}

func TestLinkedListGet(t *testing.T) {
	linked_list := linked_lists.NewLinkedList[int]()

	test_cases := []int{1, 2, 3}

	for i, v := range test_cases {
		linked_list.Append(v)

		got, _ := linked_list.Get(i)

		assert.Equal(t, got, v)
	}

	assert.Equal(t, linked_list.Length(), 3)
}
