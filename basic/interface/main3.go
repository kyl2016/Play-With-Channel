package main

import "fmt"

type Device struct {
	Name string
}

func main()  {
	ds := getDevices()

	fmt.Println(ds)

	for _, d := range ds.([]interface{}) {
		fmt.Println(d)
	}
}

func getDevices() interface{} {
	var devices []Device
	devices = append(devices, Device{"1"})
	devices = append(devices, devices...)

	var ds interface{}
	ds= devices
	return ds
}