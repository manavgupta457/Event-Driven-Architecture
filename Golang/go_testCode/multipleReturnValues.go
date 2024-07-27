package main

import (
	"fmt"
	"strings"
)

//Use an _ whenever you do not have a value for a mandatory variable

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}

func multipleReturn() {
	//Need two variable since two values are being returned
	firstName, lastName := getInitials("tifa lockhart")
	fmt.Println(firstName, lastName)

}
