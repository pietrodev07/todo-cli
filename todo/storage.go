package todo

import (
	"encoding/json"
	"io"
	"os"
)

// SaveToFile saves the task list to a JSON file.
func (t *TaskList) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(t.Items, "", "  ")
	if err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

// LoadFromFile loads the task list from a JSON file.
func (t *TaskList) LoadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File does not exist, start fresh
		}
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &t.Items)
}
