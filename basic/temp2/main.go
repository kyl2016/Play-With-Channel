package main

import (
	"log"
	"time"

	"some/path/to/lib"
)

type ClientTrace struct {
	traces       []ClientTrace
	innerOnPings []func() func(error)
	OnPing       func() func(error)
	CallOnPing   func() func(error)
}

func (c ClientTrace) Compose(in ClientTrace) ClientTrace {
	c.traces = append(c.traces, in)
	c.innerOnPings = append(c.innerOnPings, in.OnPing)
	return c
}

func (c ClientTrace) OnPing(err error) {
	for _, t := range c.innerOnPings {
		t(err)
	}
}

//
//func (c ClientTrace) OnPing(err error) {
//	for _, t := range c.traces {
//		t.Onping(err)
//	}
//}

func main() {
	var trace lib.ClientTrace // 程序 grpc等接口的后台处理，可能有错误，多种打印方法

	// Logging hooks.
	trace = trace.Compose(
		lib.ClientTrace{
			OnPing: func() func(error) {
				log.Println("ping start")
				return func(err error) {
					log.Println("ping done", err)
				}
			},
		})

	// Some metrics hooks.
	trace = trace.Compose(
		lib.ClientTrace{
			OnPing: func() func(error) {
				start := time.Now()
				return func(err error) {
					metric := stats.Get("ping_latency")
					metric.Send(time.Since(start))
				}
			},
		})

	c := lib.Client{
		Trace: trace,
	}
}
