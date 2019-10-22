package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Before(now.AddDate(0, 0, 1)))
	fmt.Println(now.Before(now.Add(time.Nanosecond)))
	fmt.Println(now.Before(now.Add(-time.Nanosecond)))
	fmt.Println(now.Before(now))
}
