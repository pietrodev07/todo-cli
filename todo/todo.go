package todo

import (
	"fmt"
)

// Task represents a single to-do task.
type Task struct {
	Title     string
	Completed bool
}

// TaskList represents a list of tasks.
type TaskList struct {
	Items []Task
}

// Add adds a new task to the list.
func (t *TaskList) Add(title string) {
	t.Items = append(t.Items, Task{Title: title, Completed: false})
}

// Remove removes a task by its index from the list.
func (t *TaskList) Remove(index int) (string, error) {
	if index < 0 || index >= len(t.Items) {
		return "", fmt.Errorf("index out of range")
	}
	task := t.Items[index].Title
	t.Items = append(t.Items[:index], t.Items[index+1:]...)
	return task, nil
}

// Complete marks a task as complete.
func (t *TaskList) Complete(index int) (string, error) {
	if index < 0 || index >= len(t.Items) {
		return "", fmt.Errorf("index out of range")
	}
	t.Items[index].Completed = true
	return t.Items[index].Title, nil
}

// List returns all tasks with their status.
func (t *TaskList) List() []string {
	var tasks []string
	for i, task := range t.Items {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		tasks = append(tasks, fmt.Sprintf("%d. %s %s", i+1, status, task.Title))
	}
	return tasks
}
