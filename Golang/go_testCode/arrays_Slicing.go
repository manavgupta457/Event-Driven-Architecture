package main

import "fmt"

func arrayNotes() {
	//Arrays need to be of fixed length
	var ages = [3]int{20, 25, 30}
	fmt.Println(ages, len(ages))
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	fmt.Println(names, len(names))

	/*
		slices (use arrays under the hood)
		The only difference being they do not have a length specified, hence we can append elements to them
	*/
	var scores = []int{100, 200, 80}
	scores = append(scores, 85)
	fmt.Println(scores)

	//slice ranges - Includes starting number but not the ending number
	//We need the shorthand only once during declaration
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
