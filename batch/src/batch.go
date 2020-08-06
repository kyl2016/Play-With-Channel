package src

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Batch struct {
	enable          bool // 是否可用
	inputCh         chan interface{}
	batchCount      int
	pollingInterval time.Duration
	batchItems      []interface{}
	ctx             context.Context
	cancel          func()
	wgPolling       sync.WaitGroup
	locker          sync.RWMutex
	stopCh          chan struct{}
	stopLocker      sync.Mutex
	batchProcessor  *Processor
}

func NewBatch(batchCount int, pollingInterval time.Duration, process func([]interface{}) error, processConcurrent int) *Batch {
	p := NewProcessor(100, processConcurrent, process)
	ctx, cancel := context.WithCancel(context.Background())
	r := Batch{
		enable:          true,
		inputCh:         make(chan interface{}, 100),
		batchCount:      batchCount,
		pollingInterval: pollingInterval,
		ctx:             ctx,
		cancel:          cancel,
		wgPolling:       sync.WaitGroup{},
		locker:          sync.RWMutex{},
		stopCh:          make(chan struct{}),
		stopLocker:      sync.Mutex{},
		batchProcessor:  p,
	}

	r.wgPolling.Add(1)
	go r.batching()
	return &r
}

// Add one item, support concurrent call
func (b *Batch) Add(item interface{}) error {
	b.locker.RLock()
	defer b.locker.RUnlock()

	if !b.enable {
		return errors.New("Adding is disable")
	}

	b.inputCh <- item
	return nil
}

// Stop the batching, support concurrent call
// after Stop, enable=false, Adding will be error
func (b *Batch) Stop() {
	b.stopLocker.Lock()
	defer b.stopLocker.Unlock()

	if !b.enable {
		return
	}
	b.stopCh <- struct{}{}
	b.wgPolling.Wait()
}

func (b *Batch) batching() {
	defer b.wgPolling.Done()

	addToBatch := func() {
		b.batchProcessor.Add(b.batchItems)
		b.batchItems = []interface{}{}
	}

	ticker := time.NewTicker(b.pollingInterval)
	for {
		select {
		case item, ok := <-b.inputCh:
			if !ok {
				b.inputCh = nil
				addToBatch()
				b.batchProcessor.Flush()
				return
			}
			b.batchItems = append(b.batchItems, item)
			if len(b.batchItems) >= b.batchCount {
				addToBatch()
			}
		case <-ticker.C:
			addToBatch()
		case <-b.stopCh:
			b.locker.Lock()
			ticker.Stop()
			b.enable = false
			close(b.inputCh)
			b.stopCh = nil
			b.locker.Unlock()
		}
	}
}
