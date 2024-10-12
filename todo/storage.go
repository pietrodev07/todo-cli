package todo

import (
	"encoding/json"
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
