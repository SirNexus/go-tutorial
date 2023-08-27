package worker

import (
	"context"
	"log"
	"sync"
)

type Worker struct {
	idx int
}

func NewWorker(idx int) *Worker {
	return &Worker{
		idx: idx,
	}
}

func (w *Worker) Run(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		w.runWorker(ctx)
		wg.Done()
	}()
}

func (w *Worker) runWorker(ctx context.Context) {
	log.Printf("worker %v started!\n", w.idx)
	for {
		select {
		case <-ctx.Done():
			w.log("finished. Exiting")
			return
		}
	}
}

func (w *Worker) log(msg string) {
	log.Printf("worker %v %v\n", w.idx, msg)
}
