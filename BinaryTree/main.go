package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

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

func (t *tree) Max() (int, error) {
	n := t.root
	if n == nil {
		return 0, errors.New("trying to get a maximum value from empty tree")
	}
	for {
		if n.right == nil {
			return n.key, nil
		}
		n = n.right
	}
}

func (t *tree) Min() (int, error) {
	n := t.root
	if n == nil {
		return 0, errors.New("trying to get a minimum value from empty tree")
	}
	for {
		if n.left == nil {
			return n.key, nil
		}
		n = n.left

	}
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

func (t *tree) PrintFromLowestToHighest() {
	defer fmt.Println()
	fmt.Printf("[ ")
	defer fmt.Printf("]")
	if t.root == nil {
		return
	} else {
		nprint(os.Stdout, t.root)
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

func nprint(writer io.Writer, n *node) {
	if n == nil {
		return
	}
	nprint(writer, n.left)
	if _, err := writer.Write([]byte(strconv.Itoa(n.key) + " ")); err != nil {
		return
	}
	//fmt.Printf("%d ", n.key) // if we don't want to test method
	nprint(writer, n.right)
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

}
