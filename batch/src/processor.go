package src

import (
	"errors"
	"sync"
)

type Processor struct {
	batchCh    chan []interface{}
	wgComplete sync.WaitGroup
	concurrent int
	process    func(items []interface{}) error
	locker     sync.RWMutex
	enable     bool
}

func NewProcessor(len, concurrent int, process func(items []interface{}) error) *Processor {
	p := Processor{make(chan []interface{}, len), sync.WaitGroup{}, concurrent, process, sync.RWMutex{}, true}
	p.wgComplete.Add(1)
	go p.startProcessing()
	return &p
}

func (p *Processor) Add(items []interface{}) error {
	p.locker.RLock()
	defer p.locker.RUnlock()

	if !p.enable {
		return errors.New("Adding is disable")
	}

	if len(items) > 0 {
		p.batchCh <- items
		items = []interface{}{}
	}

	return nil
}

func (p *Processor) Flush() {
	p.locker.Lock()
	p.locker.Unlock()

	close(p.batchCh)
	p.enable = false
	p.wgComplete.Wait()
}

func (p *Processor) startProcessing() {
	defer p.wgComplete.Done()
	//defer b.wgPolling.Done()

	wg := sync.WaitGroup{}
	for i := 0; i < p.concurrent; i++ {
		wg.Add(1)
		go func() {
			for items := range p.batchCh {
				p.process(items)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
