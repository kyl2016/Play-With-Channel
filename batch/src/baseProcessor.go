package src

import (
	"errors"
	"sync"
)

type BaseProcessor struct {
	locker  sync.RWMutex
	enable  bool
	inputCh chan interface{}
	wg      sync.WaitGroup
}

func NewBaseProcessor(concurrent int, startProcessing func(chan interface{})) *BaseProcessor {
	p := BaseProcessor{
		locker:  sync.RWMutex{},
		enable:  true,
		inputCh: make(chan interface{}, 100),
		wg:      sync.WaitGroup{},
	}

	for i := 0; i < concurrent; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			startProcessing(p.inputCh)
		}()
	}

	return &p
}

func (p *BaseProcessor) Add(item interface{}) error {
	p.locker.RLock()
	defer p.locker.RUnlock()

	if !p.enable {
		return errors.New("Add disable")
	}

	p.inputCh <- item

	return nil
}

func (p *BaseProcessor) Stop() {
	p.locker.Lock()
	p.locker.Unlock()

	if !p.enable {
		return
	}

	p.enable = false
	close(p.inputCh)
	p.wg.Wait()
}
