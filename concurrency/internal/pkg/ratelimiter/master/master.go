package master

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/bootstrap"
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/limiter"
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/worker"
)

type Master struct {
	workers []*worker.Worker

	opts *bootstrap.Options
	wg   *sync.WaitGroup

	// limiter holds the shared state across goroutines
	limiter *limiter.ThreadSafeLimiter
}

func NewMaster(opts *bootstrap.Options) *Master {
	m := &Master{
		opts:    opts,
		wg:      &sync.WaitGroup{},
		limiter: &limiter.ThreadSafeLimiter{},
	}

	for i := 0; i < m.opts.NumWorkers; i++ {
		m.addWorker(i, m.limiter)
	}

	return m
}

func (m *Master) addWorker(idx int, limiter *limiter.ThreadSafeLimiter) {
	worker := worker.NewWorker(idx, m.opts.RPS, limiter)
	m.workers = append(m.workers, worker)
}

func (m *Master) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	m.watchInterrupt(cancel)

	triggerEventChan := make(chan struct{})

	for _, worker := range m.workers {
		worker.Run(ctx, m.wg, triggerEventChan)
	}

	fmt.Println("Press enter to send a request: ")

	for {
		_, err := fmt.Scanln()
		if err != nil {
			return err
		}

		triggerEventChan <- struct{}{}
	}
}

func (m *Master) watchInterrupt(cancel context.CancelFunc) {
	cancelChan := make(chan os.Signal)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		for {
			select {
			case <-cancelChan:
				cancel()
				m.wg.Wait()
				os.Exit(0)
			}
		}
	}()
}

func (m *Master) log(msg string) {
	log.Printf("master %v\n", msg)
}
