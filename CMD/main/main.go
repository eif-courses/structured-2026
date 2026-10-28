package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func printStatus(status string) {
	fmt.Println("----------------YOUR STATUS----------------------")
	fmt.Println(status)
	fmt.Println("-------------------------------------------------")
}

func printFunction() {
	fmt.Println("IT WORKS!!!!")

}

func conditionalsStatements() {
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

var n int
var options = [][]string{
	{"1. Print your status", "printStatus"},
	{"2. Check conditional statements", "condStaments"},
	{"0. Exit", "exit"}}

var actions = map[string]func(){
	"printStatus":  printFunction,
	"condStaments": printFunction,
}

var actionNumber int

func mainblock() {
	for {
		fmt.Println("Choose one of the following options:")

		// {"1. Print your status", "printStatus"},

		for _, option := range options {
			for _, opt := range option {
				fmt.Printf(",%s ", opt)

			}
			fmt.Println()
		}

		fmt.Scan(&n)

		selectedOption := options[n-1][1]

		if fn, ok := actions[selectedOption]; ok {
			fn()
		} else {
			fmt.Println("Unknown action")
		}

	}
}

// IsDone
//

type Task struct {
	ID          int
	Description string
	Status      bool
}

// slice of tasks
// list of tasks
var TaskList []Task

var counter int = 0

func createTask(Desc string) {
	task := Task{ID: counter, Description: Desc, Status: false}
	TaskList = append(TaskList, task)
	counter++
}

// task1 , task2, task3 .. taskN
// 100000000
// CPU resources
// Nil is points to zero address in memory
func findTask(id int) *Task {
	for i := 0; i < len(TaskList); i++ {
		if TaskList[i].ID == id {
			return &TaskList[i]
		}
	}
	return nil
}

func setTaskStatus(id int, status bool) {
	task := findTask(id)
	if task != nil {
		task.Status = status
	} else {
		fmt.Println("Task not found")
	}
}

func printAllTasks() {

	for i := 0; i < len(TaskList); i++ {
		fmt.Printf("%d. %s %t \n", TaskList[i].ID, TaskList[i].Description, TaskList[i].Status)
	}
}

func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func readInt(reader *bufio.Reader) (int, error) {

	val, err := strconv.Atoi(readLine(reader))
	if err != nil {
		return 0, err // Return the error if it's not a valid number
	}

	// If you specifically want to handle a sentinel value like -1 or check something:
	if val > 10 {
		return -1, errors.New("TASK ID NOT EXIST")
	}

	return val, nil

}

type ConsoleInputVariables struct {
	Description string
	Status      int
	Input       int
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	consoleVariables := ConsoleInputVariables{Description: "", Status: 0, Input: 0}

	for {

		fmt.Println("ENTER DESCRIPTION")
		consoleVariables.Description = readLine(reader)

		createTask(consoleVariables.Description)

		printAllTasks()

		fmt.Println("ENTER TASK ID TO UPDATE STATUS")
		input, err := readInt(reader)

		if err != nil {
			log.Fatal(err)
		}

		taskID := findTask(input)

		if taskID == nil {
			fmt.Println("Task not found")
			return
		}

		fmt.Println("ENTER STATUS (1 for done, 0 for not done)")
		consoleVariables.Status, err = readInt(reader)

		if err != nil {
			log.Fatal(err)
		}

		if consoleVariables.Status == 1 {
			setTaskStatus(input, true)
		} else {
			setTaskStatus(input, false)
		}

		printAllTasks()
	}
}
