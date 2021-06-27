package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type linkedList struct{
	head *node
	tail *node
	length int
}

type node struct{
	key interface{}
	next *node
}
func newLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) Len() int{
	return l.length
}
func (l *linkedList) Display(){
	head := l.head
	fmt.Printf("[ ")
	defer fmt.Println()
	defer fmt.Printf("]")

	for i := 0; i < l.length; i++{
		fmt.Printf("%v ", head.key)
		head = head.next
	}
}

func (l *linkedList) Append(val interface{}, index ...int){
	if len(index) == 0{
		l.Append(val, l.Len())
		return
	}
	if l.head != nil && reflect.TypeOf(val).Kind() != reflect.TypeOf(l.head.key).Kind(){
		err := fmt.Sprintf("cannot use %v (type %s) as type %s in Append", val, reflect.TypeOf(val).Kind(),
			reflect.TypeOf(l.head.key).Kind())
		panic(err)
	}
	idx := index[0]
	n := &node{
		key:val,
	}
	switch{

	case l.head == nil:
		if idx != 0{
			panic("too big index")
		}
		l.head = n
		l.tail = n

	case idx == l.Len():
		l.tail.next = n
		l.tail = n

	case idx == 1:
		n.next = l.head.next
		l.head.next = n

	case idx > 1 && idx < l.Len():
		head := l.head
		for i:=0; i < idx - 1; i++{
			head = head.next
		}
		n.next = head.next
		head.next = n
	case idx < 0: panic("index can't be negative number")
	default: panic("too big index. Trying to add on [" + strconv.Itoa(idx) + "] while length of" +
		" linked list is " + strconv.Itoa(l.Len()) )
	}
	l.length++
}

func main(){
	list := newLinkedList()

	list.Append(1) //ok
	list.Append(2)//ok
	//list.Append(3.14) //panic
	list.Append(66,2) // ok
	//list.Append("55")
	list.Display()
}
