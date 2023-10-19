package main

import (
	"fmt"
	"sort"
	"strings"
)

func usingGoPackages() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "mario"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // does not alter the original string, it returns a new one

	fmt.Println("original string value =", greeting)

	// ToUpper
	// Index - find the index of

	fmt.Println(strings.Split(greeting, " "))

	// sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages) // it changes the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))
}
