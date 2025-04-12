package main

import "fmt"

func main() {
	a := person{name: "Abolfazl", age: 30}
	// talk(a) // Uncommenting this would cause a compiler error!
	talk(&a)
	a.speak()

	// REMINDER ABOUT METHOD SETS:
	// The 'speak()' method has a pointer receiver (*person).
	// This means it's in the method set of *person.
	// For a value of type 'person' (like 'a'), only methods with value receivers
	// would be in its method set. However, Go provides a convenience:
	// you can call pointer receiver methods on a value, and Go will
	// automatically take the address of the value for you (&a.speak() is valid).
	//
	// The 'speaker' interface requires a 'speak()' method.
	// - *person implements 'speaker' because it has the 'speak()' method.
	// - person does NOT directly implement 'speaker' because its 'speak()'
	//   method is defined on *person.
	//
	// Therefore, 'talk' which takes a 'speaker' can accept a *person,
	// but not a plain 'person' value directly without taking its address.

}

type person struct {
	name string
	age  int
}

type speaker interface {
	speak()
}

func talk(s speaker) {
	s.speak()
}

func (p *person) speak() {
	fmt.Printf("I'm %s and I'm %d year old\n", p.name, p.age)
}
