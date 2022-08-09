package main

import "fmt"

/*
types of linked list:
1. single linked list
2. double linked list
3. circular linked list
*/

type node struct {
	data string
	next *node
}

type mySingleLinkedList struct {
	size int
	head *node
}

func (sl *mySingleLinkedList) addToHead(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == newNode {
		fmt.Println("Cannot add (values are same as head)")
	} else {
		if sl.head == nil {
			sl.head = newNode
		} else {
			newNode.next = sl.head
			sl.head = newNode
		}
		sl.size++
	}
}

func (sl *mySingleLinkedList) addToTail(name string) {
	newNode := &node{
		data: name,
	}
	if sl.head == newNode {
		fmt.Println("Cannot add (values are same as head)")
	} else {
		if sl.head == nil {
			sl.head = newNode
		} else {
			current := sl.head
			for current.next != nil { // tail is the last node
				current = current.next
			}
			current.next = newNode
		}
		sl.size++
	}
}

func (sl *mySingleLinkedList) iterateList() {
	for node := sl.head; node != nil; node = node.next {
		fmt.Println(node.data)
	}
}

func NewLinkedList() *mySingleLinkedList {
	return &mySingleLinkedList{}
}

func main() {
	singleList := NewLinkedList()
	singleList.addToHead("A")
	singleList.addToHead("B")
	singleList.addToHead("C")
	singleList.addToTail("D")
	singleList.addToTail("E")
	singleList.iterateList()
	fmt.Println("size: ", singleList.size)
}
