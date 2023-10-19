package main

// formatting strings and printing messages
import "fmt"

// main() is the entry point of the application
func usingStrings() {
	age := 35
	name := "shaun"

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	// %v - default format specifier for variables
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	// %q - puts double quotes around strings

	fmt.Printf("age is of type %T \n", age)

	fmt.Printf("you scored %0.2f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
