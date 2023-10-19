package main

import "fmt"

func arraysSlices() {
	// arrays
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "shaun"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{100, 50, 60} // not specifying the size creates a slice
	scores[2] = 25
	scores = append(scores, 85) // returns a new slice, does not update the older one
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]  // starting from 1 but less than 3
	rangeTwo := names[2:]   // till the end
	rangeThree := names[:3] // upto index 3
	fmt.Println(rangeOne, len(rangeOne))
	fmt.Println(rangeTwo, len(rangeTwo))
	fmt.Println(rangeThree, len(rangeThree))

	rangeOne = append(rangeOne, "kappa")
	fmt.Println(rangeOne)
}
