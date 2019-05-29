package main

import (
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Local())
	fmt.Println(now.UTC())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	format := "2006-01-02 15:04:05"
	str := now.Format(format)
	fmt.Println(str)

	t, _ := time.Parse(format, str)
	fmt.Println(t)
	fmt.Println(t.Format(format))

	after := now.Add(time.Hour * 24 * 10)
	format2 := "06-01-_2"
	fmt.Println(after.Format(format2))

	f := time.Now().Format("2006-01-02")
	lastDay, _ := time.Parse("2006-01-02 15:04:05", f+" 00:00:00")
	fmt.Println(lastDay)

	maxDuration := 1<<63 - 1
	minDuration := -1 << 63
	fmt.Println("min duration:", minDuration, " max duration:", maxDuration)
	d, _ := time.ParseDuration(strconv.Itoa(maxDuration)+"ns")
	fmt.Println(d, strconv.Itoa(maxDuration))

	ctx, _ :=  context.WithTimeout(context.Background(), 1<<63-1)
	time, ok := ctx.Deadline()
	fmt.Println(time, ok)
}
