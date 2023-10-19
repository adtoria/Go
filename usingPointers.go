package main

import "fmt"

func updateName2(x *string) {
	*x = "wedge"
}

func usingPointers() {
	name := "tifa"

	//updateName2(name)

	// fmt.Println("memory address of name is: ", &name)

	ptr := &name
	fmt.Println("memory address:", ptr)
	fmt.Println("value at memory address:", *ptr)

	fmt.Println(name)
	updateName2(ptr)
	fmt.Println(name)
}
