package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go
type Node struct {
	Key   int
	Value int
	left  *Node
	right *Node
}
type OrderedMap struct {
	Root *Node
	size int
}

func NewOrderedMap() OrderedMap {
	return OrderedMap{} // need to implement
}

func (m *OrderedMap) Insert(key, value int) {
	if m.Root == nil {
		m.Root = &Node{Key: key, Value: value}
		m.size++
		return
	}
	current := m.Root
	for {
		if key < current.Key {
			if current.left == nil {
				current.left = &Node{Key: key, Value: value}
				m.size++
				return
			}
			current = current.left
		} else if key > current.Key {
			if current.right == nil {
				current.right = &Node{Key: key, Value: value}
				m.size++
				return
			}
			current = current.right
		} else {
			current.Value = value
			return
		}
	}
}

func (m *OrderedMap) Erase(key int) {
	m.Root = eraseNode(m.Root, key, &m.size)
}

func eraseNode(n *Node, key int, size *int) *Node {
	if n == nil {
		return nil
	}

	if key < n.Key {
		n.left = eraseNode(n.left, key, size)
		return n
	}

	if key > n.Key {
		n.right = eraseNode(n.right, key, size)
		return n
	}

	*size--

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	successor := minNode(n.right)

	n.Key = successor.Key
	n.Value = successor.Value

	n.right = eraseSuccessor(n.right, successor.Key)

	return n
}

func minNode(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func eraseSuccessor(n *Node, key int) *Node {
	if key < n.Key {
		n.left = eraseSuccessor(n.left, key)
	} else if key > n.Key {
		n.right = eraseSuccessor(n.right, key)
	} else {
		return n.right
	}

	return n
}

func (m *OrderedMap) Contains(key int) bool {

	for curr := m.Root; curr != nil; {
		if key == curr.Key {
			return true
		}
		if key < curr.Key {
			curr = curr.left
		} else if key > curr.Key {
			curr = curr.right
		}
	}
	return false // need to implement
}

func (m *OrderedMap) Size() int {
	return m.size // need to implement
}

func (m *OrderedMap) ForEach(action func(int, int)) {
	var walk func(*Node)

	walk = func(n *Node) {
		if n == nil {
			return
		}

		walk(n.left)
		action(n.Key, n.Value)
		walk(n.right)
	}

	walk(m.Root)
}

func TestCircularQueue(t *testing.T) {
	data := NewOrderedMap()
	assert.Zero(t, data.Size())

	data.Insert(10, 10)
	data.Insert(5, 5)
	data.Insert(15, 15)
	data.Insert(2, 2)
	data.Insert(4, 4)
	data.Insert(12, 12)
	data.Insert(14, 14)

	assert.Equal(t, 7, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(3))
	assert.False(t, data.Contains(13))

	var keys []int
	expectedKeys := []int{2, 4, 5, 10, 12, 14, 15}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))

	data.Erase(15)
	data.Erase(14)
	data.Erase(2)

	assert.Equal(t, 4, data.Size())
	assert.True(t, data.Contains(4))
	assert.True(t, data.Contains(12))
	assert.False(t, data.Contains(2))
	assert.False(t, data.Contains(14))

	keys = nil
	expectedKeys = []int{4, 5, 10, 12}
	data.ForEach(func(key, _ int) {
		keys = append(keys, key)
	})

	assert.True(t, reflect.DeepEqual(expectedKeys, keys))
}
