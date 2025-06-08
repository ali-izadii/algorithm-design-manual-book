package chapter_03

import "cmp"

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type ComparatorFunc[T any] func(a, b T) int

type Tree[T cmp.Ordered] struct {
	Root       *Node[T]
	Comparator ComparatorFunc[T]
}

func NewTree[T cmp.Ordered]() *Tree[T] {
	return &Tree[T]{
		Comparator: func(a, b T) int {
			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		},
	}
}

func newNode[T any](value T) *Node[T] {
	return &Node[T]{Value: value}
}

func (t *Tree[T]) Insert(value T) {
	t.Root = t.insert(t.Root, value)
}

func (t *Tree[T]) insert(node *Node[T], value T) *Node[T] {
	if node == nil {
		return newNode(value)
	}

	if t.Comparator(value, node.Value) < 0 {
		node.Left = t.insert(node.Left, value)
	} else if t.Comparator(value, node.Value) > 0 {
		node.Right = t.insert(node.Right, value)
	}

	return node
}

func (t *Tree[T]) Search(value T) *Node[T] {
	return t.search(t.Root, value)
}

func (t *Tree[T]) search(node *Node[T], value T) *Node[T] {
	if node == nil || t.Comparator(value, node.Value) == 0 {
		return node
	}

	if t.Comparator(value, node.Value) > 0 {
		return t.search(node.Right, value)
	}

	return t.search(node.Left, value)
}

func (t *Tree[T]) Delete(value T) {
	t.Root = t.delete(t.Root, value)
}

func (t *Tree[T]) delete(node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}

	if t.Comparator(value, node.Value) < 0 {
		node.Left = t.delete(node.Left, value)
	} else if t.Comparator(value, node.Value) > 0 {
		node.Right = t.delete(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}

		temp := t.minValueNode(node.Right)
		node.Value = temp.Value
		node.Right = t.delete(node.Right, temp.Value)
	}
	return node
}

func (t *Tree[T]) minValueNode(node *Node[T]) *Node[T] {
	current := node
	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *Tree[T]) InOrder() []T {
	var result []T
	t.inOrder(t.Root, &result)
	return result
}

func (t *Tree[T]) inOrder(node *Node[T], result *[]T) {
	if node != nil {
		t.inOrder(node.Left, result)
		*result = append(*result, node.Value)
		t.inOrder(node.Right, result)
	}
}

func (t *Tree[T]) PreOrder() []T {
	var result []T
	t.preOrder(t.Root, &result)
	return result
}

func (t *Tree[T]) preOrder(node *Node[T], result *[]T) {
	if node != nil {
		*result = append(*result, node.Value)
		t.preOrder(node.Left, result)
		t.preOrder(node.Right, result)
	}
}

func (t *Tree[T]) PostOrder() []T {
	var result []T
	t.postOrder(t.Root, &result)
	return result
}

func (t *Tree[T]) postOrder(node *Node[T], result *[]T) {
	if node != nil {
		t.postOrder(node.Left, result)
		t.postOrder(node.Right, result)
		*result = append(*result, node.Value)
	}
}
