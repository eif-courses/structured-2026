package datastructures

import (
	"fmt"
)

func RunTestForDataStructures() {

	title := "DATA STRUCTURES"
	title2 := "ARRAYS"
	title3 := "SLICES"

	// titles[0] = title, titles[1] = title2, titles[2] = title3, titles[3] = "MAPS"

	//  0        1       2      3
	// title, title2, title3, "MAPS"
	var titles [4]string
	titles[0] = title
	titles[1] = title2
	titles[2] = title3
	titles[3] = "MAPS"

	//  INCREASING INDEX by adding 1 to the index
	// i = 0, 1, 2, 3         len(titles) = 4
	// for i := 0; i < len(titles); i++ {
	// 	fmt.Println(i, " ", titles[i])
	// }

	OperationsWithSlicesAndArrays()

}

func ArrayPrinterToConsole(message string, numbers []int) {

	fmt.Println(message)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, " ", numbers[i])
	}
}

func OperationsWithSlicesAndArrays() {
	numbers := []int{1, 2, 3, 4, 5}

	ArrayPrinterToConsole("Initial numbers slice", numbers)

	numbers = append(numbers, 6, 7, 8, 9, 10)

	ArrayPrinterToConsole("After appending new numbers to the slice", numbers)

	numbers = numbers[2:4]

	ArrayPrinterToConsole("After slicing the numbers slice", numbers)
}
