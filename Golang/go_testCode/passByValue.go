package main

import "fmt"

/*
Go is a pass by value language
Go makes copies of values when passes into functions

Data types have two groups

group A types (Non-Pointer Values) -> strings, ints, bools, floats, arrays, structs
For group A data types, a copy is made of the data stored
Hence when a function is called it does not change the value of the original variable being sent to the function


group B types (Pointer Wrapper Values) -> slices, functions, map
For group B data types, a copy is made of the pointer to the data stored
Hence when a function is called it changes the value of the original variable being sent to the function
Pointers are things that point to another memory block/location
*/

//Won't cause any change to the name variable
// func updateName(x string) {
// 	x = "yoshi"
// }

// Will change the value of name variable by storing its return to the variable
func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func passByValue() {
	name := "tifa"
	name = updateName(name)
	fmt.Println(name)

	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)

}
