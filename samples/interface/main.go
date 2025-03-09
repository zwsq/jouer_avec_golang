package main

import (
	"fmt"
	"log"
	"strconv"
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

// ----------------------------------------------------------------------
type book struct {
	title string
}
type count int

func (b book) String() (selfIntroduction string) {
	return fmt.Sprint("I'm a book and my name is ", b.title)
}

func (c count) String() (howManyIAm string) {
	return fmt.Sprint("I have ", strconv.Itoa(int(c)))
}

func customeLogger(s fmt.Stringer) {
	log.Println("Log from svc", s.String())
}

// ----------------------------------------------------------------------

func main() {
	{
		// zwsq := person{
		// 	name: "Abolfazl",
		// }
		// james := secretAgent{
		// 	person: person{
		// 		name: "James",
		// 	},
		// 	ltk: true,
		// }
		// sayIt(zwsq)
		// sayIt(james)
	}
	{
		learnGo := book{
			title: "Learn Go in a week!",
		}

		var age count = 33

		fmt.Println(learnGo)
		fmt.Println(age)
		customeLogger(age)
	}
	{

	}
}
