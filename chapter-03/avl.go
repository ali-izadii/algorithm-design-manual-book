package chapter_03

import "cmp"

type AVLNode[T any] struct {
	Value  T
	Left   *AVLNode[T]
	Right  *AVLNode[T]
	Height int
}

type AVLTree[T cmp.Ordered] struct {
	Root       *AVLNode[T]
	Comparator ComparatorFunc[T]
}

func NewAVLTree[T cmp.Ordered]() *AVLTree[T] {
	return &AVLTree[T]{
		Root: nil,
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

func height[T any](n *AVLNode[T]) int {
	if n == nil {
		return 0
	}
	return n.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func newAVLNode[T any](value T) *AVLNode[T] {
	return &AVLNode[T]{Value: value, Height: 1}
}

func rightRotate[T any](y *AVLNode[T]) *AVLNode[T] {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

func leftRotate[T any](x *AVLNode[T]) *AVLNode[T] {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}

func getBalance[T any](n *AVLNode[T]) int {
	if n == nil {
		return 0
	}
	return height(n.Left) - height(n.Right)
}

func (t *AVLTree[T]) Insert(value T) {
	t.Root = t.insert(t.Root, value)
}

func (t *AVLTree[T]) insert(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return newAVLNode(value)
	}

	if t.Comparator(value, node.Value) < 0 {
		node.Left = t.insert(node.Left, value)
	} else if t.Comparator(value, node.Value) > 0 {
		node.Right = t.insert(node.Right, value)
	} else {
		return node
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))

	balance := getBalance(node)

	if balance > 1 && t.Comparator(value, node.Left.Value) < 0 {
		return rightRotate(node)
	}

	if balance < -1 && t.Comparator(value, node.Right.Value) > 0 {
		return leftRotate(node)
	}

	if balance > 1 && t.Comparator(value, node.Left.Value) > 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && t.Comparator(value, node.Right.Value) < 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

func (t *AVLTree[T]) minValueNode(node *AVLNode[T]) *AVLNode[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (t *AVLTree[T]) Delete(value T) {
	t.Root = t.delete(t.Root, value)
}

func (t *AVLTree[T]) delete(node *AVLNode[T], value T) *AVLNode[T] {
	if node == nil {
		return node
	}

	if t.Comparator(value, node.Value) < 0 {
		node.Left = t.delete(node.Left, value)
	} else if t.Comparator(value, node.Value) > 0 {
		node.Right = t.delete(node.Right, value)
	} else {
		if (node.Left == nil) || (node.Right == nil) {
			var temp *AVLNode[T]
			if node.Left != nil {
				temp = node.Left
			} else {
				temp = node.Right
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := t.minValueNode(node.Right)
			node.Value = temp.Value
			node.Right = t.delete(node.Right, temp.Value)
		}
	}

	if node == nil {
		return node
	}

	node.Height = max(height(node.Left), height(node.Right)) + 1

	balance := getBalance(node)

	if balance > 1 && getBalance(node.Left) >= 0 {
		return rightRotate(node)
	}

	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	if balance < -1 && getBalance(node.Right) <= 0 {
		return leftRotate(node)
	}

	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}
