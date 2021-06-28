package main

import "testing"

/*func TestAppend(t *testing.T){
	tests := []struct{
		name string
		val interface{}
		want interface{}
	}{
		{},
	}
}*/

func TestPop(t *testing.T){
	//Первый тест
	emptyQueue := NewLinkedList()

	//Второй тест
	queueWithNumbers := NewLinkedList()
	queueWithNumbers.tail = &node{3, nil}
	n := &node{2,queueWithNumbers.tail}
	queueWithNumbers.head = &node{1,n}
	queueWithNumbers.length += 3
	//TODO: добавить сравнение двух слайсов (comparing uncomparable type []string)
	//третий тест
	//queueWithSlices := NewLinkedList()
	//queueWithSlices.tail = &node{[]string{"dog", "cat", "elephant"}, nil}
	//queueWithSlices.head = &node{[]string{"Golang", "Python", "C++","Kotlin"}, queueWithSlices.tail}
	//queueWithSlices.length += 2

tests := []struct {
	name  string
	queue *LinkedList
	want  interface{}
	wantPanic bool
}{
	{"empty queue",emptyQueue,0, true},
	{"success pop from queue with numbers", queueWithNumbers, 3,false},
	//{"success pop from queue with slice",queueWithSlices, []string{"dog", "cat", "elephant"},false },
}
for _,test := range tests{
	t.Run(test.name, func(t *testing.T){
		defer func(){
			r := recover()
			if (r != nil) != test.wantPanic{
				t.Errorf("Pop() recover: %v, want: %v", r, test.wantPanic)
			}
		}()
		got := test.queue.Pop()
		if got != test.want{
			t.Errorf("Pop() got: %v, want: %v", got, test.want)
		}

	})
}
}
func TestLen(t *testing.T){
	//Тест 1
	emptyQueue := NewLinkedList()

	//Тест 2
	queue := NewLinkedList()
	queue.tail, queue.head = &node{2,nil}, &node{1, queue.tail}
	queue.length += 2

	tests := []struct{
		name string
		queue *LinkedList
		want interface{}
	}{
		{"success with empty queue", emptyQueue, 0},
		{"success with non empty queue",queue, 2 },
	}
	for _,test := range tests{
		t.Run(test.name, func(t *testing.T){
			got := test.queue.Len()
			if got != test.want{
				t.Errorf("%s: got: %v, want: %v", test.name, got, test.want)
			}
		})
	}
}
