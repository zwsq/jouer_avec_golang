package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	{
		user1 := user{Name: "Billie", Age: 26}
		user2 := user{Name: "Ken", Age: 65}
		users := []user{user1, user2}

		marshalled, err := json.Marshal(users)
		if err != nil {
			log.Fatal("Unable to process the object: ", err)
		}

		fmt.Println("Marshalled: ", string(marshalled))
	}
	{
		customers := []customer{}
		fmt.Println(customers)

		err := json.Unmarshal([]byte(sampleJson), &customers)
		if err != nil {
			log.Fatal("Unable to process the given json: ", err)
		}

		fmt.Println(customers)
	}
}

type user struct {
	Name string
	Age  int
}

type customer struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

var sampleJson string = `[{"firstName": "John","lastName": "Smith","age": 25}]`
