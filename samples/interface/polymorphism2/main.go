package main

import (
	"fmt"
	"log"
	"strconv"
)

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

func main() {
	{
		learnGo := book{
			title: "Learn Go in a week!",
		}

		var age count = 33

		fmt.Println(learnGo)
		fmt.Println(age)
		customeLogger(age)
	}

}
