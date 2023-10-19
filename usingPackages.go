package main

import "fmt"

var score = 99.5

func usingPackages() {

	sayHello("mario")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}
