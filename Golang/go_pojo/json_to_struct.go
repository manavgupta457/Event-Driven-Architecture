package main

import (
	"encoding/json"
	"fmt"
)

func json_to_struct1() {
	jsonString := `{"name": "John Doe", "age": 30, "email": "john.doe@example.com"}`

	var user User

	err := json.Unmarshal([]byte(jsonString), &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("User: %+v\n", user)
}

func json_to_struct2() {
	// JSON string
	jsonString := `{
        "name": "Jane Smith",
        "age": 30,
        "active": true,
        "skills": ["Go", "Python", "JavaScript"],
        "home": {
            "street": "123 Main St",
            "city": "Springfield",
            "zip": "12345"
        }
    }`

	// Create an instance of the struct to hold the unmarshaled data
	var person Person

	// Unmarshal the JSON string into the struct
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the struct
	fmt.Printf("%+v\n", person)
}
