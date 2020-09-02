package main

import (
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	f := func() {
		println("trigger")
	}

	c := cron.New()

	id, err := c.AddFunc("@every 1s", f)
	if err != nil {
		panic(err)
	}

	c.Start()

	time.Sleep(time.Second * 2)

	println("removing...")
	c.Remove(id)
	println("removed")

	time.Sleep(time.Second * 2)
}
