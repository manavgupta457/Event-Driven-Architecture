package main

import "fmt"

/*
All of the keys in a single map must have the same type
All of the values in a single map must have the same type
*/

func mapTest() {
	//In this case for menu, Keys must be string and values must be float
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 10.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//Looping Maps
	//In this case it is key, value not index, value as in case of arrays and slices

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phonebook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[845775485])

	phonebook[984759373] = "bowser"
	fmt.Println(phonebook)

	phonebook[845775485] = "yoshi"
	fmt.Println(phonebook)

}
