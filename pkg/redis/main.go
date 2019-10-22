package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	m := map[string]interface{}{
		"name1": "kitty",
		"name2": "Bob",
	}

	err = client.HMSet("test", m).Err()
	if err != nil {
		panic(err)
	}

	r, err := client.HMGet("test", "name1").Result()
	fmt.Println(r, err)

	r2, err := client.HGet("test", "name1").Result()
	fmt.Println(r2, err)

	buffer, err := client.HGet("test", "name1").Bytes()
	println("buffer len=" + strconv.Itoa(len(buffer)))

	buffer, err = client.HGet("test", "name100").Bytes()
	println("name100 buffer len=" + strconv.Itoa(len(buffer)))

	r, err = client.HMGet("test", "name100").Result()
	if err != nil {
		panic(err)
	}

	if len(r) > 0 {
		if r[0] == nil {
			fmt.Println("name100 not exist")
		}
	}

	count := client.Exists("test").Val()
	println("Exist " + strconv.Itoa(int(count)))

	client.Del("test")

	count = client.Exists("test").Val()
	println("Exist " + strconv.Itoa(int(count)))

	items := client.Keys("ivs.repos.*.people").Val()
	println(len(items))
}
