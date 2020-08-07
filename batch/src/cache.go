package src

import "sync"

type MyPool struct {
	new    func() []interface{}
	s      [][]interface{}
	locker sync.Mutex
}

func NewMyPool(new func() []interface{}) *MyPool {
	return &MyPool{new: new, locker: sync.Mutex{}}
}

func (p *MyPool) Get() []interface{} {
	p.locker.Lock()
	defer p.locker.Unlock()

	if len(p.s) > 0 {
		r := p.s[0]
		p.s = p.s[1:]
		return r
	}

	return p.new()
}

func (p *MyPool) Put(item []interface{}) {
	p.locker.Lock()
	defer p.locker.Unlock()

	p.s = append(p.s, item[:0])
}

func (p *MyPool) Clear() {
	p.s = nil
}
