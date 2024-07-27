package main

import (
	"bufio"
	"fmt" //used to format strings and print to console
	"os"

	"rsc.io/quote"
)

/*
The main file is the entry point to the application
The main function is the entry point
Use command go run . to run the entire package
Use command go run fileName.go to run a particular file
Variables and functions in the same package can be used in any file irrespective of the location of declaration
*/

// The below code serves as notes

func main() {

	//Can also do the below declaration without specifying the data type
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	var IntegerVal int
	IntegerVal = 20

	fmt.Println("Hello, World!")

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and remaining", remainingTickets)

	//Formatting
	//To print the variable
	fmt.Printf("We have total of %v tickets and remaining %v\n", conferenceTickets, remainingTickets)
	//Can also use %q to put quotes around String

	// To get the type
	fmt.Printf("We have total of %T tickets and remaining %T\n", conferenceTickets, remainingTickets)
	fmt.Printf("The type of IntegerVal is %T\n", IntegerVal)

	//Similar to declaring a variable using var - shorthand (Can't be used outside the function)
	anotherValue := "MG"
	fmt.Println(anotherValue)

	//Can use %0.2f etc.
	fmt.Printf("You scored %f points!\n", 225.55)

	//Sprintf (save formatted strings)
	//It does not print to the console, but is used when you want to store it and print later
	var str = fmt.Sprintf("my age is %v and my name is %v \n", IntegerVal, conferenceName)
	fmt.Println("the saved string is:", str)

	fmt.Println(quote.Go())

	//Types of Integers (bits and memory) Refer to https://go.dev/ref/spec#Numeric_types
	// var numOne int8 = 25
	// var numTwo int8 = -128
	// var numThree uint = 25 // Does not include negatives

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 1513123153131231125.521531354131351111
	// float64 has a higher precision than float32 and 64 is by default

	//Functions in all other files need to be run via main file as follows for the same package
	//arrayNotes()
	//standardLibrary()
	//loops()
	// boolConditionals()
	// functionsTestMain()
	// multipleReturn()
	// mapTest()
	// passByValue()
	// pointersTest()

	// myBill := newBill("mario's bill")
	// myBill.addTip(10)
	// myBill.addItem("chocolava", 1.2323153135)
	// fmt.Println(myBill.format())

	// myBill := createBill()
	// fmt.Println(myBill)

	// userInput()

	//To test switch and saving files
	// myBill := createBill()
	// promptOptions(myBill)

	//Interfaces
	// interfaces()

	//Testing Gorotines
	// goroutines()

	//Testing Channels
	// channels()

	//Testing select on channels
	selectMain()

}

/*
	Reading User Input
	We'll need to import a package bufio
	Newreader is the method in which we specify the source of the information(from where to read)
*/

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b

}
