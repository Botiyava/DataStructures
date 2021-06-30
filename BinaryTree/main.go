package main

import "fmt"

type node struct {
	key         int
	left, right *node
}

type tree struct {
	root *node
}

func NewTree() *tree {
	return &tree{}
}
func (t *tree) Append(data int) {
	insertNode := &node{
		key: data,
	}
	if t.root == nil {
		t.root = insertNode
	} else {
		t.root.insert(insertNode)
	}
}

//Печатает элементы дерева от наименьшего до наибольшего
func (t *tree) PrintFromLowestToHighest() {
	defer fmt.Println()
	fmt.Printf("[ ")
	defer fmt.Printf("]")
	if t.root == nil {
		return
	} else {
		nprint(t.root)
	}
}

func (t *tree) Exist(data int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	ok := exist(t.root, data)
	if !ok {
		return 0, false
	}
	return data, true
}

func (n *node) insert(node *node) {
	if node.key == n.key {
		fmt.Println("такой элемент уже есть")
		return
	}
	if node.key < n.key {
		if n.left == nil {
			n.left = node
		} else {
			n.left.insert(node)
		}
	} else {
		if n.right == nil {
			n.right = node
		} else {
			n.right.insert(node)
		}
	}
}

func nprint(n *node) {
	if n == nil {
		return
	}
	nprint(n.left)
	fmt.Printf("%d ", n.key)
	nprint(n.right)
}

func exist(root *node, data int) (ok bool) {
	queue := make([]*node, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.key == data {
			return true
		}
		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}
	return false
}
func main() {
	t := NewTree()
	t.Append(8)
	t.Append(4)
	t.Append(7)
	t.Append(3)
	t.PrintFromLowestToHighest()
	if _, ok := t.Exist(7); !ok {
		fmt.Println("Такого элемента нет")
	}

}
