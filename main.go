package main

import (
	"fmt"
)

func printStatus(status string) {
	fmt.Println("----------------YOUR STATUS----------------------")
	fmt.Println(status)
	fmt.Println("-------------------------------------------------")
}

func main() {

	status := "IN PROGRESS"
	done := true
	isReady := true

	if isReady && !done {
		fmt.Println("Preparing your order")
		status = "PREPARING"
		printStatus(status)

		printStatus("Order is READY")
		done = true
		if done {
			printStatus("DONE order delivered")
		}
	} else {
		fmt.Println("ALL ORDERS IS COMPLETED!!!")
	}

	// i = 1   1 < 5
	// i = 2  2 < 5
	// i = 3 3< 5
	// i = 4 4<5

	// i = 5    5 < 5

	//i=0,  0 % 2 == 0
	//i=1, 0 % 2 == 1

	for i := 0; i < 5; i = i + 1 {

		if i%2 == 0 {
			fmt.Println(i+1, "EVEN NUMBER ", i)
		} else {
			fmt.Println(i+1, "ODD NUMBER", i)
		}
	}

	numbers := []int{231, 222, 2222, 1111, 43}

	for i := 0; i < len(numbers); i++ {

		//fmt.Println(numbers[i])
		temp := numbers[2:4]

		for j := 0; j < len(temp); j++ {
			fmt.Println(temp[j])
		}
		break

	}

}
