package src

import (
	"context"
	"sync"
	"time"
)

type Batch struct {
	enable          bool // 是否可用
	inputCh         chan interface{}
	batchCount      int
	pollingInterval time.Duration
	process         func([]interface{}) error
	batchItems      []interface{}
	ctx             context.Context
	cancel          func()
	wgPolling       sync.WaitGroup
}

func NewBatch(batchCount int, pollingInterval time.Duration, process func([]interface{}) error) *Batch {
	ctx, cancel := context.WithCancel(context.Background())
	r := Batch{
		enable:          true,
		inputCh:         make(chan interface{}, 100),
		batchCount:      batchCount,
		pollingInterval: pollingInterval,
		process:         process,
		ctx:             ctx,
		cancel:          cancel,
		wgPolling:       sync.WaitGroup{},
	}

	r.wgPolling.Add(1)
	go r.polling()
	return &r
}

func (b *Batch) Add(item interface{}) error {
	b.inputCh <- item
	return nil
}

func (b *Batch) Stop() {
	b.cancel()
	b.wgPolling.Wait()
}

func (b *Batch) polling() {
	defer b.wgPolling.Done()

	ticker := time.NewTicker(b.pollingInterval)
	for {
		select {
		case item, ok := <-b.inputCh:
			if !ok {
				return
			}
			b.batchItems = append(b.batchItems, item)
			if len(b.batchItems) >= b.batchCount {
				b.process(b.batchItems)
			}
		case <-ticker.C:
			b.process(b.batchItems)
		case <-b.ctx.Done():
			close(b.inputCh)
			return
		}
	}
}
