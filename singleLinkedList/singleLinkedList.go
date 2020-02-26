package main

import (
	"fmt"
)

// SingleNode - of a linked list
type SingleNode struct {
	value int
	nextNode *SingleNode
}

// SetValue - sets value of a node
func (node *SingleNode) SetValue (value int) {
	node.value = value
}

// GetValue - returns the value of node
func (node *SingleNode) GetValue () int {
	return node.value
}

// NewSingleNode - creates & returns a new empty node
func NewSingleNode() *SingleNode {
	return new(SingleNode)
}

// SingleLinkedList - defining the structure
type SingleLinkedList struct {
	head *SingleNode
	tail *SingleNode
}

// AddNode - adds a new node to single linked list
func (list *SingleLinkedList) AddNode(value int) {
	newSingleNode := NewSingleNode()
	newSingleNode.SetValue(value)

	if list.head == nil {
		list.head = newSingleNode
	} else if list.head == list.tail {
		list.head.nextNode = newSingleNode
	} else if list.tail != nil {
		list.tail.nextNode = newSingleNode
	}
	list.tail = newSingleNode
}

// GetAllValues - returns the values from list
func (list *SingleLinkedList) GetAllValues() []int {
	var values []int
	tempNode := list.head
	for tempNode != nil {
		values = append(values, tempNode.GetValue())
		tempNode = tempNode.nextNode
	}
	return values
}

func (list *SingleLinkedList) String() string {
	var s string
	for node := list.head; node != nil; node = node.nextNode {
		s = s + fmt.Sprintf(" {%d} ->", node.GetValue())
	}
	return s + " nil"
}

// NewSingleLinkedList - creates a new Single Linked List
func NewSingleLinkedList() *SingleLinkedList {
	return new(SingleLinkedList)
}

func main() {
	myList := NewSingleLinkedList()
	myList.AddNode(2)
	myList.AddNode(20)
	myList.AddNode(200)
	myList.AddNode(2000)
	myList.AddNode(20000)
	fmt.Println("List: ", myList.GetAllValues())
	fmt.Println("Stringer List: ", myList.String())
}