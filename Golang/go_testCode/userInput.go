package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Different methods for user Input

fmt.Scan reads space-separated values from the standard input.

fmt.Scanln reads space-separated values and stops scanning at a newline.

fmt.Scanf reads input according to a format specifier.

For reading an entire line of input, including spaces, use the bufio package.

*/

func userInput() {
	var name string
	var age int

	fmt.Println("Enter your name and age:")
	fmt.Scanf("%s %d", &name, &age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	fmt.Println("Enter your name and age:")
	//Stops scanning after a space
	fmt.Scan(&name)
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	fmt.Println("Enter your name and age:")
	//Stops scanning after a space
	fmt.Scanln(&name)
	fmt.Scanln(&age)

	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

	// //If you want to input a full name then that is possible only with bufio.NewReader
	// //It ensures that the input captures everything the user types until they press enter

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)

}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

//Function to save the bill

func (b bill) save() {
	data := []byte(b.format())

	//We'll be using the os package to save a file
	//In the parentheses enter the file location
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}
