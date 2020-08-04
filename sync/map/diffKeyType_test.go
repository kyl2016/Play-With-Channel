package main

import (
	"sync"
	"testing"
)

func TestMap_DiffKeyType(t *testing.T) {
	m := sync.Map{}
	var i int = 1
	m.Store(i, i)
	var j int32 = 1
	_, ok := m.Load(j)
	if ok {
		t.Error("can't use int32 as key of type int")
	}
}
