package src

import (
	"context"
	"sync"
	"time"
)

type Batch struct {
	enable            bool // 是否可用
	inputCh           chan interface{}
	batchCount        int
	pollingInterval   time.Duration
	process           func([]interface{}) error
	batchItems        []interface{}
	ctx               context.Context
	cancel            func()
	wgPolling         sync.WaitGroup
	processConcurrent int
}

func NewBatch(batchCount int, pollingInterval time.Duration, process func([]interface{}) error, processConcurrent int) *Batch {
	ctx, cancel := context.WithCancel(context.Background())
	r := Batch{
		enable:            true,
		inputCh:           make(chan interface{}, 100),
		batchCount:        batchCount,
		pollingInterval:   pollingInterval,
		process:           process,
		ctx:               ctx,
		cancel:            cancel,
		wgPolling:         sync.WaitGroup{},
		processConcurrent: processConcurrent,
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
	batchCh := make(chan []interface{}, 100)
	addToBatch := func() {
		if len(b.batchItems) > 0 {
			batchCh <- b.batchItems
			b.batchItems = []interface{}{}
		}
	}

	go func() {
		defer b.wgPolling.Done()

		wg := sync.WaitGroup{}
		for i := 0; i < b.processConcurrent; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for items := range batchCh {
					b.process(items)
				}
			}()
		}
		wg.Wait()
	}()

	stopped := false
	ticker := time.NewTicker(b.pollingInterval)
	for {
		select {
		case item, ok := <-b.inputCh:
			if !ok {
				b.inputCh = nil
				addToBatch()
				close(batchCh)
				return
			}
			b.batchItems = append(b.batchItems, item)
			if len(b.batchItems) >= b.batchCount {
				addToBatch()
			}
		case <-ticker.C:
			addToBatch()
		case <-b.ctx.Done():
			if !stopped {
				close(b.inputCh)
				stopped = true
			}
		}
	}
}
