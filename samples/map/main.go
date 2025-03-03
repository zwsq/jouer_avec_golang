package main

import (
	"fmt"
)

func main() {
	iMap := map[string]string{"home_range": "192.168.2.0/24", "company_range": "192.168.1.0/24", "lab_range": "192.168.102.0/24"}
	nMap := make(map[int]int)

	// Looping over map keys and values
	for k, v := range iMap {
		fmt.Printf("K: %v\tV: %v\n", k, v)
	}

	// Looping over map keys
	fmt.Println("-----------------------------------------------------")
	for k := range iMap {
		fmt.Println(k)
	}

	// Looping over map values
	fmt.Println("-----------------------------------------------------")
	for _, k := range iMap {
		fmt.Println(k)
	}

	// Assigning a values to maps
	fmt.Println("-----------------------------------------------------")
	for i := 1; i <= 10; i++ {
		nMap[i] = 10 * i
	}
	fmt.Println(nMap)

	// Deleting item from map
	fmt.Println("-----------------------------------------------------")
	delete(iMap, "company_range") // if the key does not exist, the compiler won't panic.
	fmt.Println(iMap)

	// Making sure the item really exists since the default value would return in case of
	// accessing a key which does not exist
	if v, ok := iMap["Akbar"]; !ok {
		fmt.Println("Queried value does not exist.")
	} else {
		fmt.Printf("exits %v\n", v)
	}
	exercises()

}
func exercises() {
	fmt.Println("************************************************************")
	fmt.Println("************************************************************")
	fmt.Println("************************************************************")

	map49 := make(map[string][]string)
	map49[`bond_jame`] = []string{`Say my name`, `what's your name`, `she's a here`}
	map49[`zwsq`] = []string{`danger`, `road runner`, `migmig`}

	for k, v := range map49 {
		fmt.Printf("values for the key: %v ==> ", k)
		for i, v2 := range v {
			fmt.Printf("%v %v ", i, v2)
		}
		fmt.Println()
	}
}
