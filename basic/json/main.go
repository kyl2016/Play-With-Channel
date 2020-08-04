package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// 检查json是否按之前的顺序进行序列化

func main() {
	t1, _ := time.Parse("2006-01-02 15:04:05.999999", "0001-01-01 00:00:17.240000")
fmt.Printf("%v\n", t1)
	t, _ := time.Parse(time.RFC3339, "0001-01-01T08:06:00.24+08:05")
	//t, _=time.ParseDuration("16.6s")
	fmt.Printf("%v\n", t)
	r, _ := json.Marshal(t)
	fmt.Println(string(r))

	return

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	expect := "[1,2,3,4,5,6,7,8,9]"

	for {
		buffer, _ := json.Marshal(nums)
		if string(buffer) != expect {
			panic(string(buffer))
		}
	}

}
