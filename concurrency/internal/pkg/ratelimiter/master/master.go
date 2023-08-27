package master

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/bootstrap"
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/worker"
)

type Master struct {
	workers []*worker.Worker

	opts *bootstrap.Options
	wg   *sync.WaitGroup
}

func NewMaster(opts *bootstrap.Options) *Master {
	m := &Master{
		opts: opts,
		wg:   &sync.WaitGroup{},
	}

	for i := 0; i < m.opts.NumWorkers; i++ {
		m.addWorker(i)
	}

	return m
}

func (m *Master) addWorker(idx int) {
	worker := worker.NewWorker(idx)
	m.workers = append(m.workers, worker)
}

func (m *Master) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	m.watchInterrupt(cancel)

	for _, worker := range m.workers {
		worker.Run(ctx, m.wg)
	}

	fmt.Println("Press enter to send a request: ")

	for {
		_, err := fmt.Scanln()
		if err != nil {
			return err
		}
		fmt.Println("Great job!")
	}
}

func (m *Master) watchInterrupt(cancel context.CancelFunc) {
	cancelChan := make(chan os.Signal)
	signal.Notify(cancelChan, os.Interrupt)

	go func() {
		for {
			select {
			case <-cancelChan:
				log.Println("Cancel interrupt found")
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
