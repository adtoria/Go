package main

import "fmt"

func usingMaps() {
	// maps in go
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		267584976: "mario",
		438924925: "luigi",
		874389182: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[267584976])

	phonebook[267584976] = "shaun"

	fmt.Println(phonebook[267584976])
}
