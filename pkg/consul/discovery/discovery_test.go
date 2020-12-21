package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestCallService(t *testing.T) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		panic(err)
	}

	//svc, _ := connect.NewService("hello", client)
	//defer svc.Close()
	//httpClient := svc.HTTPClient()
	//
	//resp, err := httpClient.Get("https://userinfo.service.consul/user/mitchellh")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(resp)

	//
	c := Client{Consul: client}
	//	err = c.Register("myService", 5039)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	entry, meta, err := c.Service("hello", "primary")
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Consul.
		fmt.Println(entry, meta)
}
