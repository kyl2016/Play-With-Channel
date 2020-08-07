package src

import (
	"time"
)

type BatchImp struct {
	batchProcessor *BaseProcessor
	batching       *BaseProcessor
}

func NewBatchImp(batchCount int, pollingInterval time.Duration, processConcurrent int, process func(interface{}) error) *BatchImp {
	p := NewMyPool(func() []interface{} {
		return []interface{}{}
	})

	batchProcessor := NewBaseProcessor(processConcurrent, startProcessBatchDatas(process, p))

	r := &BatchImp{
		batchProcessor: batchProcessor,
		batching:       NewBaseProcessor(1, startBatching(pollingInterval, batchCount, batchProcessor, p)),
	}

	return r
}

func (b *BatchImp) Add(item interface{}) error {
	return b.batching.Add(item)
}

func (b *BatchImp) Stop() {
	b.batching.Stop()
	b.batchProcessor.Stop()
}

func startProcessBatchDatas(process func(interface{}) error, p *MyPool) func(inputCh chan interface{}) {
	return func(inputCh chan interface{}) {
		for items := range inputCh {
			process(items)
			p.Put(items.([]interface{}))
		}
	}
}

func startBatching(pollingInterval time.Duration, batchCount int, processor *BaseProcessor, p *MyPool) func(chan interface{}) {
	var batchItems = p.Get()

	return func(inputCh chan interface{}) {
		addToBatch := func() {
			if len(batchItems) > 0 {
				processor.Add(batchItems)
				batchItems = p.Get()
			}
		}

		ticker := time.NewTicker(pollingInterval)
		for {
			select {
			case item, ok := <-inputCh:
				if !ok {
					addToBatch()
					processor.Stop()
					return
				}
				batchItems = append(batchItems, item)
				if len(batchItems) >= batchCount {
					addToBatch()
				}
			case <-ticker.C:
				addToBatch()
			}
		}
	}
}
