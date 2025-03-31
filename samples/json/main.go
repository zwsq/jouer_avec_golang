package main

import (
	"encoding/json"
	"fmt"
	"log"
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
func main() {
	c1 := customer{Name: "Abolfazl", Cell: "552313455", Address: "513, California city blvd, California City, California, USA", Zip: "979565", Age: 30}
	c2 := customer{Name: "Sajad", Cell: "423664234", Address: "423, Ontario Hall, British Clombia, Canada", Zip: "38958KF", Age: 31}
	customers := []customer{
		c1,
		c2,
	}
	j, err := json.Marshal(customers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))

	unmarshalledCustomers := []customer{}
	err = json.Unmarshal(j, &unmarshalledCustomers)
	if err != nil {
		log.Fatal("Error unmarshaling json file ", err)
	}

	fmt.Println(unmarshalledCustomers)
}
