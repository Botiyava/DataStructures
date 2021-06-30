package main

import (
	"testing"
)

func TestAppend(t *testing.T){
	//First test
	t1 := NewTree()
	//Second test
	t2 := NewTree()
	t2.root = &node{
		key:8,
	}
	t2.root.left = &node{
		key: 6,
		left: &node{
			key: 4,
		},
		right: &node{
			key: 7,
		},
	}
	t2.root.right = &node{
		key: 24,
	}

	tests := []struct{
		name string
		tree *tree
		data int
		wantBeforeAppend bool
		wantAfterAppend bool
		//TODO: if element > than left and < than right

	}{
		{"empty tree", t1, 4,false,true},
		{"append in random tree", t2, 5, false, true},
		{"element already exist", t2, 7, true,true},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			_, existBefore := test.tree.Exist(test.data)
			test.tree.Append(test.data)
			_, existAfter := test.tree.Exist(test.data)
			if existBefore != test.wantBeforeAppend {
				t.Errorf("Append(): Exist(): got: %v, want: %v", existBefore, test.wantBeforeAppend)
			}
			if existAfter != test.wantAfterAppend {
				t.Errorf("Append(): Exist(): got: %v, want: %v", existAfter, test.wantAfterAppend)
			}

		})
	}
}
func TestMax(t *testing.T){
	//First test
	t1 := NewTree()
	//Second test
	t2 := NewTree()
	t2.Append(4)
	//Third test
	t3 := NewTree()
	t3.Append(16)
	t3.Append(8)
	t3.Append(24)
	t3.Append(12)

	tests := []struct{
		name string
		tree *tree
		wantInt int
		wantErr bool
	}{
		{"empty tree", t1, 0, true},
		{"maximum - root node", t2, 4, false},
		{"maximum - random node", t3,  24, false},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			gotInt, gotErr := test.tree.Max()
			if gotInt != test.wantInt{
				t.Errorf("Max(): got: %v, want: %v", gotInt, test.wantInt)
			}
			if (gotErr != nil) != test.wantErr{
				t.Errorf("Max(): got: %v, want: %v", gotErr, test.wantErr)
			}
		})
	}
}

func TestMin(t *testing.T){
	//First test
	t1 := NewTree()
	//Second test
	t2 := NewTree()
	t2.Append(4)
	//Third test
	t3 := NewTree()
	t3.Append(16)
	t3.Append(8)
	t3.Append(24)
	t3.Append(12)

	tests := []struct{
		name string
		tree *tree
		wantInt int
		wantErr bool
	}{
		{"empty tree",t1, 0,true},
		{"minimum = root node", t2, 4, false},
		{"find minimum in random place", t3, 8, false},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			gotInt, gotErr := test.tree.Min()
			if gotInt != test.wantInt{
				t.Errorf("Min(): got: %v, want: %v", gotInt, test.wantInt)
			}
			if (gotErr != nil) != test.wantErr{
				t.Errorf("Min(): got: %v, want: %v", gotErr, test.wantErr)
			}
		})
	}
}
func TestExist(t *testing.T){
	//First test
	t1 := NewTree()
	//Second test
	t2 := NewTree()
	t2.Append(4)
	//Third test
	t3 := NewTree()
	t3.Append(16)
	t3.Append(8)
	t3.Append(24)
	t3.Append(12)


	tests := []struct{
		name string
		data int
		tree *tree
		wantInt int
		wantBool bool
	}{
		{"empty tree", 4,t1,0, false },
		{"find root node", 4, t2, 4, true},
		{"find node in random place", 12, t3, 12, true},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){

			gotInt, gotBool := test.tree.Exist(test.data)
			if gotInt != test.wantInt{
				t.Errorf("NewTree(): got: %v, want: %v", gotInt, test.wantInt)
			}
			if gotBool != test.wantBool{
				t.Errorf("NewTree(): got: %v, want: %v", gotBool, test.wantBool)
			}
		})
	}
}
