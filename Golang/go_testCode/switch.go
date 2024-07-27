package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		//Parsing a float - Here we are converting the string input to a float
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("item added -", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.addTip(t)
		fmt.Println("Tip has been updated to", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("The bill has been saved as", b.name)

	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}
