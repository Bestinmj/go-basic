package main

import "fmt"

func main() {
	anotherVar := "maybe"
	var test int = 2
	fmt.Println(anotherVar)
	fmt.Println("test!")
	fmt.Println(test)
	var message = fmt.Sprintf("testing, %v. Working", test)
	fmt.Println(message)

	var userInput string
	fmt.Println("enter some string: ")
	fmt.Scan(&userInput)
	fmt.Printf("user input: %v", userInput)

}
