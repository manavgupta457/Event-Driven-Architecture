package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Refer to this link https://pkg.go.dev/std
Import any functionality you need and start using it
*/

func standardLibrary() {
	greeting := "hello there hello friends!"
	fmt.Println(strings.Contains(greeting, "hello!"))        //To check whether it contains hello! in greeting variable
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) //Does not modify the original string instead returns a new string
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Split(greeting, " ")) //To split by a particular value

	var splitGreet = strings.Split(greeting, " ")
	fmt.Println(splitGreet[0])

	//the original value is unchanged
	fmt.Println("Original :", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) //Sorts the slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) //Returns the index
	fmt.Println(index)

	fmt.Println(sort.SearchInts(ages, 90)) //Returns length+1 since value does not exist and higher than the highest
	fmt.Println(sort.SearchInts(ages, 10)) //Returns 0 since value does not exist and lower than the lowest
	fmt.Println(sort.SearchInts(ages, 40)) //Returns the position where it should ideally be added to maintain the order

	names := []string{"yoshi", "mario", "peach", "bowser"}
	sort.Strings(names) //Sort by alphabetical order
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "zeta")) //Same logic as SearchInts

}
