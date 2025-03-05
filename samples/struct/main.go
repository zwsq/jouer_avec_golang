package main

import (
	"fmt"
)

type options struct {
	gpu   bool
	tenG  bool
	varez bool
}

type vritualMachine struct {
	ip         string
	hostname   string
	os         string
	env        string
	cpuCores   int
	memoryInGB float32
	vlan       string
	options    options // embedded struct
}

type person struct {
	firstName      string
	lastName       string
	favoriteColors []string
}

func initialExercises() {

	projectX_Dev := vritualMachine{
		ip:         "192.168.1.200",
		hostname:   "projectx.dev.local",
		os:         "debian 12",
		env:        "dev",
		cpuCores:   24,
		memoryInGB: 128.8,
		vlan:       "enterprise solutions",
		// embedded struct
		options: options{
			gpu:   false,
			tenG:  true,
			varez: true,
		},
	}
	projectX_Prod := vritualMachine{
		ip:         "192.168.100.100",
		cpuCores:   48,
		memoryInGB: 512,
		vlan:       "Protected",
		options: options{
			gpu:  true,
			tenG: true,
		},
	}

	// Anonymous struct: It is defined as the variable is being defined.
	human := struct {
		name    string
		speaker bool
		brain   bool
	}{
		name:    "Daniel",
		speaker: true,
		brain:   true,
	}

	fmt.Printf("%#v\n", projectX_Dev)
	fmt.Printf("%#v\n", projectX_Prod)
	fmt.Printf("%#v\n", human)
}

// Struct with slice
func exercise54() {
	girl := person{
		firstName:      "Jenny",
		lastName:       "Juano",
		favoriteColors: []string{"Pink", "Purple"},
	}
	boy := person{
		firstName:      "James",
		lastName:       "Bond",
		favoriteColors: []string{"Blue", "Green"},
	}

	for _, v := range girl.favoriteColors {
		fmt.Println(girl.firstName, girl.lastName, "loves", v)
	}
	for _, v := range boy.favoriteColors {
		fmt.Println(boy.firstName, boy.lastName, "loves", v)
	}
}

// Struct with map
func exercise55() {
	girl := person{
		firstName:      "Jenny",
		lastName:       "Juano",
		favoriteColors: []string{"Pink", "Purple"},
	}
	boy := person{
		firstName:      "James",
		lastName:       "Bond",
		favoriteColors: []string{"Blue", "Green"},
	}

	staffBook := map[string]person{
		girl.lastName: girl,
		boy.lastName:  boy,
	}
	for _, v := range staffBook {
		fmt.Println(v)
		for _, v2 := range v.favoriteColors {
			fmt.Println(v.firstName, "loves", v2)
		}
	}
}

// Struct with embedded struct
func exercise56() {
	type engine struct {
		electric bool
	}
	type car struct {
		engine
		brand string
		doors int
		color string
	}

	love := car{
		brand:  "Volkswagen Golf GTi",
		color:  "Red",
		doors:  4,
		engine: engine{false},
	}
	hatred := car{
		brand:  "BMW *",
		color:  "White",
		doors:  4,
		engine: engine{true},
	}
	fmt.Println(love)
	fmt.Println(hatred)
}

// Struct with anonymous struct
func exercise57() {

}

func main() {

	exercise56()

}
