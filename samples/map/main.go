package main

import "fmt"

func main() {
	iMap := map[string]string{"home_range": "192.168.2.0/24", "company_range": "192.168.1.0/24", "lab_range": "192.168.102.0/24"}
	nMap := make(map[int]int)

	// Looping over map
	for k, v := range iMap {
		fmt.Printf("K: %v\tV: %v\n", k, v)
	}

	for i := 1; i <= 10; i++ {
		nMap[i] = 10 ^ i
	}

	fmt.Println(nMap)
}
