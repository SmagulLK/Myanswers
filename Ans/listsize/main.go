package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func main() {
	link := &List{}
	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")
	fmt.Println(ListSize(link))
}

func ListPushFront(l *List, data interface{}) {
	value := &NodeL{Data: data}
	if l.Tail == nil {
		l.Tail = value
		l.Head = value
	} else {
		value.Next = l.Head
		l.Head = value
	}
}

func ListSize(l *List) int {
	count := 0
	temp := l.Head
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}
