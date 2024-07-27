package main

import "fmt"

// here * is the dereference operator - Used to retrieve the value from the memory address
func updatePointerName(n *string) {
	*n = "wedge"
}

func pointersTest() {
	name := "tifa"
	m := &name //to get address of name variable

	fmt.Println("memory address: ", m)
	fmt.Println("value at memory address: ", *m)
	updatePointerName(m)
	fmt.Println(name)

}
