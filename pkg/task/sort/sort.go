package sort

import (
	"fmt"
	"goqueue/pkg/array"
	"time"
)

// SortTask holds the attributes needed to perform unit of work
type SortTask struct {
	ID       int
	arr      array.IArray
	writerID int
	createDt int64
}

// New creates SortTask object, takes an array for sorting and a producer id
func New(arr array.IArray, writerID int) *SortTask {
	return &SortTask{
		arr:      arr,
		writerID: writerID,
		createDt: time.Now().Unix(),
	}
}

// Identity return ID of task
func (t *SortTask) Identity() int {
	return t.ID
}

// Exec execute task, returns an error if present
func (t *SortTask) Exec() error {
	fmt.Printf("%d [%s] %d %d %d\n",
		t.writerID,
		time.Unix(t.createDt, 0).Format(time.RFC3339),
		t.arr.Min(),
		t.arr.Median(),
		t.arr.Max(),
	)
	return nil
}
