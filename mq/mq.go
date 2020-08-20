package mq

import (
	"github.com/google/uuid"
	"sync"
)

// MQ light mq with channel
type MQ struct {
	typeAndChans sync.Map // type and map[chan interface{}]struct{} mapping
	idAndCh      sync.Map // id and chan interface{} mapping
	locker       sync.RWMutex
}

func NewMQ() *MQ {
	return &MQ{typeAndChans: sync.Map{}, idAndCh: sync.Map{}, locker: sync.RWMutex{}}
}

// Sub multi types data to channel with id
func (mq *MQ) Sub(types ...string) (id string, c <-chan interface{}) {
	mq.locker.Lock()
	defer mq.locker.Unlock()

	id := uuid.New().String()
	ch := make(chan interface{}, 100)
	mq.idAndCh.Store(id, ch)
	for _, t := range types {
		mChsI, ok := mq.typeAndChans.Load(t)
		if ok {
			mChs := mChsI.(sync.Map)
			mChs.Store(ch, struct{}{})
		} else {
			mChs := sync.Map{} // map[chan interface{}]struct{}{}
			mChs.Store(ch, struct{}{})
			mq.typeAndChans.Store(t, mChs)
		}
	}

	return id, ch
}

// Unsub by id and types
func (mq *MQ) Unsub(id string, types ...string) {
	mq.locker.Lock()
	defer mq.locker.Unlock()

	chI, ok := mq.idAndCh.Load(id)
	if !ok {
		return
	}
	ch := chI.(chan interface{})

	for _, t := range types {
		mI, ok := mq.typeAndChans.Load(t)
		if !ok {
			continue
		}
		m := mI.(sync.Map)
		_, exist := m.Load(ch)
		if exist {
			m.Delete(ch)
			close(ch)
		}
	}
}

// Pub data with type
func (mq *MQ) Pub(t string, data interface{}) {
	mq.locker.RLock()
	defer mq.locker.RUnlock()

	mChsI, ok := mq.typeAndChans.Load(t)
	if !ok {
		return
	}
	mChs := mChsI.(sync.Map)
	mChs.Range(func(chI, _ interface{}) bool {
		ch := chI.(chan interface{})
		ch <- data
		return true
	})
}
