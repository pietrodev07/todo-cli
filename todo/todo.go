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
