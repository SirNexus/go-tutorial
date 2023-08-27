package worker

import (
	"log"
)

type Worker struct {
	idx int
}

func NewWorker(idx int) *Worker {
	return &Worker{
		idx: idx,
	}
}

func (w *Worker) Run() {
	go w.runWorker()
}

func (w *Worker) runWorker() {
	log.Printf("worker %v started!\n", w.idx)
}
