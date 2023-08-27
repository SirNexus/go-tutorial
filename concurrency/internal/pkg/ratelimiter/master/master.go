package master

import (
	"fmt"

	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/bootstrap"
	"github.com/SirNexus/go-tutorial/concurrency/internal/pkg/ratelimiter/worker"
)

type Master struct {
	workers []*worker.Worker

	opts *bootstrap.Options
}

func NewMaster(opts *bootstrap.Options) *Master {
	m := &Master{
		opts: opts,
	}

	for i := 0; i < opts.NumWorkers; i++ {
		m.addWorker(i)
	}

	return m
}

func (m *Master) addWorker(idx int) {
	worker := worker.NewWorker(idx)
	m.workers = append(m.workers, worker)
}

func (m *Master) Run() error {
	fmt.Println("Press enter to send a request: ")

	for {
		_, err := fmt.Scanln()
		if err != nil {
			return err
		}
		fmt.Println("Great job!")
	}
}
