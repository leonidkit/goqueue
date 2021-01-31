package writer

import (
	"goqueue/internal/queue"

	intarr "goqueue/pkg/array/intarr"
	tasksort "goqueue/pkg/task/sort"

	// taskmock "goqueue/pkg/task/mock"
	"sync"
)

// Writer is a producer of tasks for queue
type Writer struct {
	writersCount int
	iterCount    int
	arrSize      int
}

// New returns Writer object, takes writers count, iteration count
func New(writersC, iterC, arrSize int) *Writer {
	return &Writer{
		writersCount: writersC,
		iterCount:    iterC,
		arrSize:      arrSize,
	}
}

// Produce run starts writers putting their tasks in a queue, takes the queue to be put in
func (w *Writer) Produce(q *queue.QueueTasks) {
	go func() {
		var wg sync.WaitGroup
		wg.Add(w.writersCount)

		for wr := 0; wr < w.writersCount; wr++ {
			go w.enqueue(&wg, q, wr)
		}

		wg.Wait()
		q.Delete()
	}()
}

func (w *Writer) enqueue(wg *sync.WaitGroup, q *queue.QueueTasks, writerID int) {
	defer wg.Done()
	for i := 0; i < w.iterCount; i++ {
		arr := intarr.New(intarr.GenerateRandomSlice(w.arrSize))
		q.Push(tasksort.New(arr, writerID))
	}
}
