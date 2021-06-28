package main

import (
	"errors"
	"fmt"
)

type slice struct{
	elements []int
}
func (slice *slice) get(index int) (int, error){
	if len(slice.elements) < index || len(slice.elements) == 0{
		return 0, errors.New("элемента с таким индексом не существует или массив пустой")
	}
	return slice.elements[index], nil
}

func (slice *slice) insert( index, value int) error{
	if len(slice.elements) < index{
		return errors.New("сюда нельзя вставить элемент")
	}
	if len(slice.elements) == 0 || len(slice.elements) == index{
		slice.elements = append(slice.elements, value)
		return  nil
	}
	slice.elements = append(slice.elements[:index+ 1], slice.elements[index:]...)
	slice.elements[index] = value
	return  nil
}

func (slice *slice) delete(index int) error{
	if len(slice.elements) < index || len(slice.elements) == 0{
		return errors.New("элемента с таким индексом не существует или он пустой")
	}
	slice.elements = append(slice.elements[:index], slice.elements[index + 1:]...)
	return nil
}

func (slice *slice) size() int{
	return len(slice.elements)
}
func (slice *slice) reverse(){
	for i := 0; i < len(slice.elements)/2; i++{
		slice.elements[i], slice.elements[len(slice.elements) - i - 1] = slice.elements[len(slice.elements) - i - 1], slice.elements[i]
	}
}
func (slice *slice) theSecondBiggest() int{
	max := 0
	secondMax := 0
	for _, val := range slice.elements{
		if val > max{
			secondMax = max
			max = val
			continue
		}else{
			if val > secondMax{
				secondMax = val
			}

		}
	}
	return secondMax
}
func (slice *slice) sort(){
	//TODO: дописать функцию сортировки массива
}
func main(){
	arr := &slice{[]int{1,2,3,4,5,6}}
	arr.insert(2, 99)
	fmt.Println(arr)
	fmt.Println(arr.get(2))
	arr.delete(2)
	fmt.Println(arr)
	fmt.Println("The size of arr is:", arr.size())
	fmt.Println(arr.theSecondBiggest())
	arr.reverse()
	fmt.Println(arr.elements)
}
