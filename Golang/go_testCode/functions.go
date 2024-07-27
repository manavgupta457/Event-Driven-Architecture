package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// When we want to return something, we need to mention the data type
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func functionsTestMain() {
	sayGreeting("User")
	sayBye("User")

	//Here we'll run sayGreeting for each value in slice
	cycleNames([]string{"cloud", "jira", "bill"}, sayGreeting)

	//Area of a circle
	a1 := circleArea(10.5)
	fmt.Printf("Area of circle 1 is %0.3f", a1)

}
