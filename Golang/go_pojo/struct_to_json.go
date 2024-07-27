package main

import (
	"encoding/json"
	"fmt"
)

func struct_to_json1() {
	// Create an instance of the struct with values
	user := User{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	// Convert the struct to JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}

// Trying with different data types
// Define a nested struct
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	Zip    string `json:"zip"`
}

// Define the main struct with different data types
type Person struct {
	FullName string   `json:"name"`   // String type
	Years    int      `json:"age"`    // Integer type
	IsActive bool     `json:"active"` // Boolean type
	Skills   []string `json:"skills"` // Slice of strings
	Home     Address  `json:"home"`   // Nested struct
}

func struct_to_json2() {
	// Create an instance of the struct with values
	person := Person{
		FullName: "Jane Smith",
		Years:    30,
		IsActive: true,
		Skills:   []string{"Go", "Python", "JavaScript"},
		Home: Address{
			Street: "123 Main St",
			City:   "Springfield",
			Zip:    "12345",
		},
	}

	// Convert the struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the JSON string
	fmt.Println(string(jsonData))
}
