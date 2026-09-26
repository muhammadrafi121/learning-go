package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var tasks []string

	greetings := "\n==================================\n" +
		"1. Add new task\n" +
		"2. List all tasks\n" +
		"3. Edit a task\n" +
		"4. Remove task by index\n" +
		"5. Exit\n" +
		"==================================\n" +
		"Your Choice: "

	var choice int

	for {
		fmt.Print(greetings)
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		switch choice {
		case 1:
			fmt.Print("New task: ")

			if scanner.Scan() {
				task := scanner.Text()
				tasks = append(tasks, task)
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error:", err)
				return
			}

		case 2:
			fmt.Println("List of tasks:")
			for index, value := range tasks {
				fmt.Printf("%d. %s\n", index+1, value)
			}
		case 3:
			var editID int
			fmt.Print("Task index: ")
			_, err := fmt.Scanln(&editID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			fmt.Print("New task: ")

			if scanner.Scan() {
				task := scanner.Text()
				tasks[editID-1] = task
			}

			if err := scanner.Err(); err != nil {
				fmt.Println("Error:", err)
				return
			}
		case 4:
			var removeID int
			fmt.Print("Task index: ")
			_, err := fmt.Scanln(&removeID)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			tasks = append(tasks[:removeID-1], tasks[removeID:]...)
		case 5:
			fmt.Println("See ya! :D")
			return
		}
	}
}
