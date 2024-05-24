package main

import (
	"fmt"
	"strings"
)

func test(firstName string) string {
	fmt.Printf("seconf function has happened %v\n", firstName)
	return "testing"
}

func main() {
	var testArray = [50]string{}
	fmt.Println(testArray)
	var names = []string{}
	var firstName = ""
	var secondName = ""

	fmt.Println("Enter the first name")
	fmt.Scan(&firstName)
	fmt.Println("Enter the second name")
	fmt.Scan(&secondName)

	returnValue := test(firstName)
	fmt.Printf("return value %v\n", returnValue)

	isValidName := len(firstName) >= 2 && len(secondName) >= 2

	if !isValidName {
		fmt.Println("Wrong input")
	}
	names = append(names, firstName+" "+secondName)
	for _, element := range names {
		var test = strings.Fields(element) // slice each element
		println(test[0])
	}
}
