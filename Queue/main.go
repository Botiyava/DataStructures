package main

import (
	"fmt"
	"reflect"
)

type LinkedList struct {
	head   *node
	tail   *node
	length int
}

type node struct {
	key  interface{}
	next *node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Len() int {
	return l.length
}
func (l *LinkedList) Display() {
	head := l.head
	fmt.Printf("[ ")
	defer fmt.Println()
	defer fmt.Printf("]")

	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ", head.key)
		head = head.next
	}
}
func (l *LinkedList) Append(val interface{}) {

	if l.head != nil && reflect.TypeOf(val).Kind() != reflect.TypeOf(l.head.key).Kind() {
		panic(fmt.Sprintf("cannot use %v (type %s) as type %s in Append", val, reflect.TypeOf(val).Kind(),
			reflect.TypeOf(l.head.key).Kind()))
	}
	n := &node{
		key: val,
		next: l.head,
	}
	l.head = n
	l.length++
}

func (l *LinkedList) Pop() interface{}{
	if l.head == nil {
		panic("nothing to delete from this queue")
	}
	head := l.head
	for i := 0; i < l.Len() - 2; i++{
		head = head.next
	}
	item := head.next
	head.next = nil
	l.length -= 1
	return item.key
}
func main() {

	}

