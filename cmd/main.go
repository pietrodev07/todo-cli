package main

import (
	"fmt"
	"os"
	"strconv"
	"todo-cli/todo"
)

const todoFile = "todo.json"

func main() {
	list := &todo.TaskList{}

	if err := list.LoadFromFile(todoFile); err != nil {
		fmt.Println("Error loading the list:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: [add|remove|complete|list] [task/index]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		task := os.Args[2:]

		if len(task) == 0 {
			fmt.Println("Specify a task")
			return
		}

		list.Add(task[0])
		fmt.Printf("Task added: %s\n", task[0])

	case "remove":
		index, _ := strconv.Atoi(os.Args[2])
		task, err := list.Remove(index - 1)

		if err != nil {
			fmt.Println("Something went wrongs:", err)
			return
		}

		fmt.Printf("Task removed: %s\n", task)

	case "complete":
		index, _ := strconv.Atoi(os.Args[2])
		task, err := list.Complete(index - 1)

		if err != nil {
			fmt.Println("Something went wrongs:", err)
			return
		}

		fmt.Printf("Task completed: %s\n", task)

	case "list":
		tasks := list.List()

		for _, task := range tasks {
			fmt.Println(task)
		}

	default:
		fmt.Println("Unknown command:", command)
	}

	if err := list.SaveToFile(todoFile); err != nil {
		fmt.Println("Error saving the list:", err)
	}
}
