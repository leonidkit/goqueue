package main

import (
	"flag"
	"goqueue/internal/queue"
	"goqueue/internal/writer"
	"log"
)

var (
	writersCount int
	iterCount    int
	arrSize      int
)

func main() {
	flag.IntVar(&writersCount, "writers", 1, "количество пишущих горутин")
	flag.IntVar(&iterCount, "iter-count", 10, "количество итераций")
	flag.IntVar(&arrSize, "arr-size", 1_000_000, "количество итераций, максимум 9223372036854775807")

	flag.Parse()

	q := queue.New(writersCount)

	writers := writer.New(
		writersCount,
		iterCount,
		arrSize,
	)
	writers.Produce(q)

	// main горутина выступает в роли исполнителя
	for {
		task := q.Pop()
		if task == nil {
			break
		}

		err := task.Exec()
		if err != nil {
			log.Printf("task ID = %d failed: %v", task.Identity(), err)
		}
	}
}
