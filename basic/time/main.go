package main

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	stat := runtime.MemStats{}
	runtime.ReadMemStats(&stat)
	fmt.Printf("%v", stat)

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Day())
	fmt.Println(now.YearDay())
	fmt.Println(now.Local())
	fmt.Println(now.UTC())
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	date1 := "2019-09-04T15:59:59Z"
	date, _ := time.Parse(time.RFC3339, date1)
	fmt.Println(date)
	fmt.Println(date.Format("2006-03-04 05:06:07.999"))

	format := "2006-01-02 15:04:05"
	str := now.Format(format)
	fmt.Println(str)

	t, _ := time.Parse(format, str)
	fmt.Println(t)
	fmt.Println(t.Format(format))

	after := now.Add(time.Hour * 24 * 10)
	format2 := "06-01-_2"
	fmt.Println(after.Format(format2))
	fmt.Println(now.Format(format2))

	format2 = "06-01-02"
	fmt.Println(after.Format(format2))
	fmt.Println(now.Format(format2))

	after3, err := time.Parse("2006-01-02", "2019-01-01")
	if err != nil {
		panic(err)
	}
	format3 := "06-01-02"
	fmt.Println(after3.Format(format3))

	f := time.Now().Format("2006-01-02")
	lastDay, _ := time.Parse("2006-01-02 15:04:05", f+" 00:00:00")
	fmt.Println(lastDay)

	maxDuration := 1<<63 - 1
	minDuration := -1 << 63
	fmt.Println("min duration:", minDuration, " max duration:", maxDuration)
	d, _ := time.ParseDuration(strconv.Itoa(maxDuration) + "ns")
	fmt.Println(d, strconv.Itoa(maxDuration))

	ctx, _ := context.WithTimeout(context.Background(), 1<<63-1)
	time, ok := ctx.Deadline()
	fmt.Println(time, ok)
}
