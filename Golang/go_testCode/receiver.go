package main

import "fmt"

/*
Here we'll look at receiver functions
A receiver function can be called only via the type specified
*/

//We'll be using bill struct type and newBill function from structs.go file

// Format the bill
func (b bill) format() string {
	fs := "Bill Breakdown:\n"
	var total float64 = 0

	//list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	//add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)
	total += b.tip

	//Here "total:" acts as a variable for %v
	//-25 means this variable will take up 25 characters on the right and positive means on the left

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "total:", total)
	return fs
}

//Now we'll use receiver functions with pointers

// update tip
// func (b bill) addTip(tip float64) {
// 	b.tip = tip
// }

/*
The above function won't be able to update the tip value
because in receiver functions it creates a copy and it will update the copy
Hence to update values with receiver functions we'll need to use pointers
*/

/*
When we are working with structs, we do not need to use a dereference operator for memory address
because go dereferences pointers for structs automatically
We just need to pass in the pointer to the function
So, wherever you are updating a value through  function, use pointers
The benefit with pointers is that you are not making a copy of the data each time you call a function
You are just making a copy of the pointer
*/

func (b *bill) addTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
