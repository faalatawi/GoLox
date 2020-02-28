package main

import "fmt"

type Visitor interface {
	// visitForParent(*Parent)
	visitForChildA(*ChildA)
	visitForChildB(*ChildB)
}

type Element interface {
	getType() string
	accept(Visitor)
}

// type Parent struct {
// }

type ChildA struct {
	name   string
	friend Element
}

func (c *ChildA) getType() string {
	return "ChildA"
}

func (c *ChildA) accept(v Visitor) {
	v.visitForChildA(c)
}

type ChildB struct {
	name   string
	friend Element
}

func (c *ChildB) getType() string {
	return "ChildB"
}

func (c *ChildB) accept(v Visitor) {
	v.visitForChildB(c)
}

// The visitor

type printVisitor struct {
	// empty
}

func (p printVisitor) visitForChildA(c *ChildA) {
	fmt.Println("the child:", c.getType())
}

func (p printVisitor) visitForChildB(c *ChildB) {
	fmt.Println("the child:", c.getType())
}

func main() {
	myA := &ChildA{name: "aaa"}
	myB := &ChildB{name: "bbb"}

	myV := &printVisitor{}

	myA.accept(myV)
	myB.accept(myV)

}
