package worker

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/limiter"
)

type Worker struct {
	idx     int
	rps     float32
	limiter *limiter.ThreadSafeLimiter
}

func NewWorker(idx int, rps float32, limiter *limiter.ThreadSafeLimiter) *Worker {
	return &Worker{
		idx:     idx,
		rps:     rps,
		limiter: limiter,
	}
}

func (w *Worker) Run(ctx context.Context, wg *sync.WaitGroup, triggerEventChan <-chan struct{}) {
	wg.Add(1)
	go func() {
		w.runWorker(ctx, triggerEventChan)
		wg.Done()
	}()
}

func (w *Worker) runWorker(ctx context.Context, triggerEventChan <-chan struct{}) {
	log.Printf("worker %v started!\n", w.idx)
	for {
		select {
		case <-triggerEventChan:
			w.processEvent()
		case <-ctx.Done():
			w.log("finished. Exiting")
			return
		}
	}
}

func (w *Worker) processEvent() {
	defer w.limiter.Unlock()

	curTime := time.Now()
	w.limiter.Lock()
	lastPacketTime := w.limiter.GetLastPacketTime()
	if curTime.Sub(lastPacketTime) < time.Duration(w.rps)*time.Second {
		w.log("event rate limited!")
		return
	}

	w.log("event passed successfully")
	w.limiter.UpdateLastPacketTime(curTime)
}

func (w *Worker) log(msg string) {
	log.Printf("worker %v %v\n", w.idx, msg)
}
