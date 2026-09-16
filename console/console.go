package console

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) PrintNameAndAge() {
	fmt.Println("Hello", p.Name)
	fmt.Println("You are", p.Age, "years old")
}

func InitMyConsoleApp() {

	var peter Person

	fmt.Println("Enter your name:")
	fmt.Scan(&peter.Name)

	fmt.Println("Enter your age:")
	fmt.Scan(&peter.Age)

	peter.PrintNameAndAge()

	fmt.Printf("value:%d address age:%p, address name:%p", peter.Age, &peter.Age, &peter.Name)
}
