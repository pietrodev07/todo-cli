package todo

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
