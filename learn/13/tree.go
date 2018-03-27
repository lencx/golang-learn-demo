package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {

		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *Tree, ch chan int) {
	rangeTree(t, ch)
	close(ch)
}

//遍历二叉树，把当前节点值传入通道
func rangeTree(t *Tree, ch chan int) {
	if t != nil {
		rangeTree(t.Left, ch)
		ch <- t.Value
		rangeTree(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *Tree) bool {
	//建立两个通道
	ch1 := make(chan int)
	ch2 := make(chan int)
	//遍历两个二叉树，把值传入各自的通道
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	//遍历通道进行比较，有不同的值就返回false
	for i := range ch1 {
		if i != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("二叉树遍历比较")
	fmt.Println("打印 New(1)的值")
	//打印 New(1)的值
	var ch1 = make(chan int)
	go Walk(New(1), ch1)
	for v := range ch1 {
		fmt.Println(v)
	}
	fmt.Println("打印 New(2)的值")
	//打印 New(2)的值
	var ch2 = make(chan int)
	go Walk(New(2), ch2)
	for v := range ch2 {
		fmt.Println(v)
	}
	//比较两个tree的值是否相等
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
}

// https://studygolang.com/articles/10837?fr=sidebar
