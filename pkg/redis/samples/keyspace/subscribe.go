package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

func main() {
	c := redis.NewClient(&redis.Options{
		DB:   0,
		Addr: "localhost:6379",
	})

	c.ConfigSet("notify-keyspace-events", "Kg") // notify 所有 keyspace 的 expired, del events

	ch := c.PSubscribe("__keyspace@0__:test*").Channel() // 订阅所有 test* 的 events
	go func() {
		for item := range ch {
			fmt.Println(strings.Replace(item.Channel, "__keyspace@0__:", "", -1), item.Payload)
			fmt.Println()
		}
	}()

	for i := 0; i < 10; i++ {
		c.Set("test"+strconv.Itoa(i), "test", time.Second*2)
	}

	for i := 0; i < 10; i++ {
		c.Set("test"+strconv.Itoa(i), "test", time.Second)
		c.Del("test" + strconv.Itoa(i))
	}

	time.Sleep(time.Second * 1)
}
