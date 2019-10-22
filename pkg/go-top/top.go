package main

import (
	"fmt"
	top "github.com/dgryski/go-topk"
)

func main(){
	 s := top.New(2)

	 es := s.Insert("1", 1)
	fmt.Println(es.Key, es.Count)

	 es = s.Insert("2.8", 1)
	fmt.Println(es.Key, es.Count)

es =	 s.Insert("3.2", 1)
	fmt.Println(es.Key, es.Count)

es =	 s.Insert("3.1", 1)
	fmt.Println(es.Key, es.Count)

es =	 s.Insert("3.0", 1)
	fmt.Println(es.Key, es.Count)

es = 	s.Insert("2", 1)
	fmt.Println(es.Key, es.Count)

	 for _, key := range s.Keys() {
	 	fmt.Println(key.Key, key.Count)
	 }
}
