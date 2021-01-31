package mock

import (
	"fmt"
)

// MockTask holds the attributes needed to perform unit of work
type MockTask struct {
	id       int
	writerID int
}

// New creates MockTask object, takes a producer id
func New(taskID, writerID int) *MockTask {
	return &MockTask{
		id:       taskID,
		writerID: writerID,
	}
}

// Identity return ID of task
func (t *MockTask) Identity() int {
	return t.id
}

// Exec execute task, returns an error if present
func (t *MockTask) Exec() error {
	fmt.Printf("task id = %d by writer id = %d done\n", t.id, t.writerID)
	return nil
}
