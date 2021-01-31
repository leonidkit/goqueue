package queue

import (
	"goqueue/pkg/task"
)

// QueueTasks is a queue of tasks
type QueueTasks struct {
	// The specification clearly says: "Channels act as first-in-first-out queues"
	jobs chan task.ITask
}

// New returns QueueTasks object, takes writers count to create a buffered channel (more efficient)
func New(writers int) *QueueTasks {
	return &QueueTasks{
		jobs: make(chan task.ITask, writers),
	}
}

// Push add task to queue
func (q *QueueTasks) Push(task task.ITask) {
	q.jobs <- task
}

// Pop return task from queue
func (q *QueueTasks) Pop() task.ITask {
	if task, ok := <-q.jobs; ok {
		return task
	}
	return nil
}

// Delete close QueueTasks jobs channel
func (q *QueueTasks) Delete() {
	close(q.jobs)
}
