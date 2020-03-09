package main

import (
	"fmt"
)

type genericNode interface {
	setValue(msg string)
	getValue() string
}

func createGenericNode() genericNode {
	gn := newSingleNode()
	gn.setValue("this is generic node - singleNode")
	return gn
}

type singleNode struct {
	nextNode *singleNode
	msg string
}

func (node *singleNode) setValue(msg string) {
	node.msg = msg
}

func (node *singleNode) getValue() string {
	return node.msg
}

func newSingleNode() *singleNode {
	return new(singleNode)
}

type singlePowerNode struct {
	nextNode *singlePowerNode
	msg string
}

func (node *singlePowerNode) setValue(msg string) {
	node.msg = msg
}

func (node *singlePowerNode) getValue() string {
	return node.msg
}

func newSinglePowerNode() *singlePowerNode {
	return new(singlePowerNode)
}

func main() {

	mySingleNode := newSingleNode()
	mySingleNode.setValue("This is normal node")
	fmt.Println("Message: ", mySingleNode.getValue())

	mySinglePowerNode := newSinglePowerNode()
	mySinglePowerNode.setValue("This is power node")
	fmt.Println("Message: ", mySinglePowerNode.getValue())

	gn := createGenericNode()

	switch concreteNode := gn.(type) {
	case *singleNode:
		fmt.Println("SingleNode Type")
		fmt.Println("Msg: ", concreteNode.getValue())
	case *singlePowerNode:
		fmt.Println("PowerNode Type")
		fmt.Println("Msg: ", concreteNode.getValue())
	}

}