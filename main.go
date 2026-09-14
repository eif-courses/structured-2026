package main

import "fmt"

var age uint = 20

func printNameAndAge(name string, age uint) {
	fmt.Println("Hello", name)
	fmt.Println("You are", age, "years old")
}

func main() {

	var name string
	var age uint

	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Println("Enter your age:")
	fmt.Scan(&age)

	printNameAndAge(name, age)

}
