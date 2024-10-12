package todo

import (
	"os"
	"testing"
)

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	taskList := &TaskList{}
	taskList.Add("Task to Save")

	tempFile, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	if err := taskList.SaveToFile(tempFile.Name()); err != nil {
		t.Fatalf("failed to save task list: %v", err)
	}

	loadedTaskList := &TaskList{}
	if err := loadedTaskList.LoadFromFile(tempFile.Name()); err != nil {
		t.Fatalf("failed to load task list: %v", err)
	}

	if len(loadedTaskList.Items) != 1 {
		t.Errorf("expected 1 task, got %d", len(loadedTaskList.Items))
	}

	if loadedTaskList.Items[0].Title != "Task to Save" {
		t.Errorf("expected task title 'Task to Save', got %s", loadedTaskList.Items[0].Title)
	}
}

func TestLoadFromNonExistentFile(t *testing.T) {
	taskList := &TaskList{}
	err := taskList.LoadFromFile("nonexistent.json")

	if err != nil {
		t.Errorf("expected no error when loading a nonexistent file, got %v", err)
	}

	if len(taskList.Items) != 0 {
		t.Errorf("expected no tasks in an empty list, got %d", len(taskList.Items))
	}
}
