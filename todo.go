package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	IsDone      bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type TaskList []item

func (tl *TaskList) AddTask(task string) {
	item := item{
		Task:        task,
		IsDone:      false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*tl = append(*tl, item)
}

func (tl *TaskList) CompleteTask(i int) error {
	l := *tl
	taskListCount := len(l)
	if i < 0 || i > taskListCount {
		return fmt.Errorf("invalid index provided")
	}

	l[i-1].IsDone = true
	l[i-1].CompletedAt = time.Now()

	return nil
}

func (tl *TaskList) SaveToFile(filename string) error {
	json, err := json.Marshal(tl)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}

func (tl *TaskList) RetrieveFromFile(fileName string) error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		// if file does not exists
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, tl)
}
