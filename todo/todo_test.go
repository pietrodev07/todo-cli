package todo

import (
	"testing"
)

func TestAdd(t *testing.T) {
	taskList := TaskList{}

	taskList.Add("Test Task 1")
	if len(taskList.Items) != 1 {
		t.Errorf("expected 1 task, got %d", len(taskList.Items))
	}

	if taskList.Items[0].Title != "Test Task 1" {
		t.Errorf("expected 'Test Task 1', got '%s'", taskList.Items[0].Title)
	}
}

func TestRemove(t *testing.T) {
	taskList := TaskList{}
	taskList.Add("Test Task 1")
	taskList.Add("Test Task 2")

	removedTask, err := taskList.Remove(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if removedTask != "Test Task 1" {
		t.Errorf("expected 'Test Task 1' to be removed, got '%s'", removedTask)
	}

	if len(taskList.Items) != 1 {
		t.Errorf("expected 1 task, got %d", len(taskList.Items))
	}

	// Test for invalid index
	_, err = taskList.Remove(5)
	if err == nil {
		t.Errorf("expected an error for invalid index, got nil")
	}
}

func TestComplete(t *testing.T) {
	taskList := TaskList{}
	taskList.Add("Test Task 1")

	// Mark the task as complete
	completedTask, err := taskList.Complete(0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if completedTask != "Test Task 1" {
		t.Errorf("expected 'Test Task 1' to be completed, got '%s'", completedTask)
	}

	if !taskList.Items[0].Completed {
		t.Errorf("expected task to be completed, got incomplete task")
	}

	// Test for invalid index
	_, err = taskList.Complete(5)
	if err == nil {
		t.Errorf("expected an error for invalid index, got nil")
	}
}

func TestList(t *testing.T) {
	taskList := TaskList{}
	taskList.Add("Test Task 1")
	taskList.Add("Test Task 2")

	tasks := taskList.List()

	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks in the list, got %d", len(tasks))
	}

	expectedFirst := "1. [ ] Test Task 1"
	if tasks[0] != expectedFirst {
		t.Errorf("expected '%s', got '%s'", expectedFirst, tasks[0])
	}

	// Mark a task as complete
	taskList.Complete(0)
	tasks = taskList.List()
	expectedFirstCompleted := "1. [x] Test Task 1"
	if tasks[0] != expectedFirstCompleted {
		t.Errorf("expected '%s', got '%s'", expectedFirstCompleted, tasks[0])
	}
}
