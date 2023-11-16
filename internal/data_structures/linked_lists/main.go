package linked_lists

import (
	"github.com/pkg/errors"
)

type ILinkedList[T comparable] interface {
	Length() int
	InsertAt(item T, index int) error
	Remove(item T) (T, error)
	RemoveAt(index int) (T, error)
	Append(item T) error
	Prepend(item T) error
	Get(index int) (T, error)
}

type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Length returns the number of elements in the linked list
func (list *LinkedList[T]) Length() int {
	count := 0
	currentNode := list.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}

// InsertAt inserts an item at the specified index in the linked list
func (list *LinkedList[T]) InsertAt(item T, index int) error {
	if index < 0 || index > list.Length() {
		return errors.Errorf("Index out of bounds")
	}

	node := &Node[T]{Data: item, Next: nil}

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		return nil
	}
	currentNode := list.Head

	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	node.Next = currentNode.Next
	currentNode.Next = node

	return nil
}

// Remove removes the first occurrence of the specified item from the linked list
func (list *LinkedList[T]) Remove(item T) (T, error) {
	var zero_value T

	if list.Head == nil {
		return zero_value, errors.Errorf("There is no linked list to delete dawg")
	}

	if list.Head.Data == item {
		removedItem := list.Head.Data
		list.Head = list.Head.Next
		return removedItem, nil
	}

	currentNode := list.Head
	for currentNode.Next != nil && currentNode.Next.Data != item {
		currentNode = currentNode.Next
	}

	if currentNode.Next == nil {
		return zero_value, errors.Errorf("The node don't exists")
	}

	removedItem := currentNode.Next.Data
	currentNode.Next = currentNode.Next.Next

	return removedItem, nil
}

// RemoveAt removes the item at the specified index from the linked list
func (list *LinkedList[T]) RemoveAt(index int) (T, error) {
	var zero_value T
	if index < 0 || index >= list.Length() {
		return zero_value, errors.Errorf("Index out of bounds")
	}

	if index == 0 {
		removedItem := list.Head.Data
		list.Head = list.Head.Next
		return removedItem, nil
	}

	currentNode := list.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	removedItem := currentNode.Next.Data
	currentNode.Next = currentNode.Next.Next

	return removedItem, nil
}

// Append adds an item to the end of the linked list
func (list *LinkedList[T]) Append(item T) error {
	node := &Node[T]{Data: item, Next: nil}

	if list.Head == nil {
		list.Head = node
		return errors.Errorf("Linked list is empty")
	}

	lastNode := list.Head

	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = node

	return nil
}

// Prepend adds an item to the beginning of the linked list
func (list *LinkedList[T]) Prepend(item T) error {
	node := &Node[T]{Data: item, Next: list.Head}
	list.Head = node

	return nil
}

func (list *LinkedList[T]) Get(index int) (T, error) {
	var zero_value T
	if index < 0 || index >= list.Length() {
		return zero_value, errors.Errorf("Index out of bounds")
	}

	currentNode := list.Head

	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}

	return currentNode.Data, nil
}
