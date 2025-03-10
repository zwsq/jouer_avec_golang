package main

import (
	"fmt"
)

type person struct {
	name string
}

type secretAgent struct {
	person person
	ltk    bool
}

type human interface {
	sayYourName()
}

func (p person) sayYourName() {
	fmt.Println("My name is", p.name)
}

func (sa secretAgent) sayYourName() {
	fmt.Println("My name is", sa.person.name, "and I'm a secrent agent.")
}

func sayIt(h human) {
	h.sayYourName()
}

func main() {
	{
		zwsq := person{
			name: "Abolfazl",
		}
		james := secretAgent{
			person: person{
				name: "James",
			},
			ltk: true,
		}
		sayIt(zwsq)
		sayIt(james)
	}
}
