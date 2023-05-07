package main

import "fmt"

// Interfaces
// An interface is a set of method signatures.
// A type implements an interface by implementing its methods.
// There is no explicit declaration of intent, no "implements" keyword.

type Human interface {
	breath()
	walk()
}

// Human interface defines a single method, breath.

// Adult and child types implement the Human interface.
type Adult struct {
	name string
}
 
func (a Adult) breath(){
	println(a.name, "IS BREATHING")

}

func (a Adult) walk() {
	println(a.name, "is walking")
}

func beingHuman(h Human) {
	h.breath()
	fmt.Println("I am being human")

}

type Child struct {
	name string
	age int
}

func (c Child) breath(){
	println("Child", c.name, "is breathing")
}
// Function that takes an interface as an argument

func main() {

}
