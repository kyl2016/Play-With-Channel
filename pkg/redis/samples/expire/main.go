package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"sync"
	"time"
)

var locker = sync.Mutex{}

func main() {
	c := redis.NewClient(&redis.Options{
		DB:   0,
		Addr: "localhost:6379",
	})

	c.ConfigSet("notify-keyspace-events", "Ex")

	keyPrefix := "position_"

	ch := c.PSubscribe("__keyevent@0__:expired").Channel()
	go monitor(c, ch)

	go func() {
		for i := 0; i < 10; i++ {
			receiveData(c, keyPrefix+strconv.Itoa(i))
		}
	}()

	time.Sleep(time.Second * 5)
}

func receiveData(c *redis.Client, key string) {
	locker.Lock()
	defer locker.Unlock()

	// check if exists
	m := c.HGetAll(key).Val()
	if len(m) == 0 {
		now := time.Now()
		c.HMSet(key, map[string]interface{}{"firstTime": now, "lastTime": now}).Err()
	} else {
		c.HSet(key, "lastTime", time.Now())
	}

	// update expire
	c.Set(key+"_expired", "", time.Second*1)
}

func monitor(c *redis.Client, ch <-chan *redis.Message) {
	for item := range ch {
		expire(c, strings.Replace(item.Payload, "_expired", "", -1))
	}
}

func expire(c *redis.Client, key string) {
	// filter

	locker.Lock()
	defer locker.Unlock()

	//m := c.HGetAll(key).Val()
	// if time.Now() - lastTime < 15s { return }

	// process

	// remove key
	fmt.Println("delete ", key, c.Del(key).Val())
}
