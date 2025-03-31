package main

import (
	"fmt"
	"sort"
)

type customer struct {
	Name    string
	Cell    string
	Address string
	Zip     string
	Age     int
}
type ByAge []customer

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []customer

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	c1 := customer{Name: "Abolfazl", Cell: "552313455", Address: "513, California city blvd, California City, California, USA", Zip: "979565", Age: 34}
	c2 := customer{Name: "Sajad", Cell: "423664234", Address: "423, Ontario Hall, British Clombia, Canada", Zip: "38958KF", Age: 31}
	customers := []customer{
		c1,
		c2,
	}
	fmt.Println("Unsorted: ", customers)
	sort.Sort(ByAge(customers))
	fmt.Println("Sorted By Age: ", customers)
	sort.Sort(ByName(customers))
	fmt.Println("Sorted By Name: ", customers)

}

func sortReporter(xs []int) {
	if sort.IntsAreSorted(xs) {
		fmt.Println("The given slice of ints is sorted.")
	} else {
		fmt.Println("The given slice of ints is not sorted. ")
	}
}
