package main

import "fmt"

func loops() {

	//While loop
	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}

	//For loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//To cycle through a slice
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	//If we only need a value and likewise for index
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

}
