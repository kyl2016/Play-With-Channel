package main

import (
	"fmt"
	"sync"
)

func main() {
	size := 10

	text1 := make(chan string, size)
	text2 := make(chan string, size)
	text3 := make(chan string, size)

	map1 := make(chan map[string]int, size)
	map2 := make(chan map[string]int, size)
	map3 := make(chan map[string]int, size)

	reduce1 := make(chan int, size)
	reduce2 := make(chan int, size)

	avg1 := make(chan float32, size)
	avg2 := make(chan float32, size)

	go inputReader([3]chan<- string{text1, text2, text3})
	go mapper(text1, map1)
	go mapper(text2, map2)
	go mapper(text3, map3)
	go shuffler([]<-chan map[string]int{map1, map2, map3}, [2]chan<- int{reduce1, reduce2})
	go reducer(reduce1, avg1)
	go reducer(reduce2, avg2)

	outputWriter([]<-chan float32{avg1, avg2})
}

func mapper(in <-chan string, out chan<- map[string]int) {
	count := map[string]int{}
	for word := range in {
		count[word] = count[word] + 1
	}
	out <- count
	close(out)
}

func reducer(in <-chan int, out chan<- float32) {
	sum, count := 0, 0
	for n := range in {
		sum += n
		count++
	}
	out <- float32(sum) / float32(count)
	close(out)
}

func inputReader(out [3]chan<- string) {
	input := [][]string{
		{"noun", "verb", "verb", "noun", "noun"},
		{"noun", "verb", "verb", "noun", "noun", "verb"},
		{"noun", "verb", "verb", "noun", "noun", "noun", "noun", "noun"},
	}

	for i := range out {
		go func(ch chan<- string, word []string) {
			for _, w := range word {
				ch <- w
			}
			close(ch)
		}(out[i], input[i])
	}
}

func shuffler(in []<-chan map[string]int, out [2]chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {
		go func(c <-chan map[string]int) {
			for m := range c {
				nc, ok := m["noun"]
				if ok {
					out[0] <- nc
				}
				vc, ok := m["verb"]
				if ok {
					out[1] <- vc
				}
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out[0])
		close(out[1])
	}()
}

func outputWriter(in []<-chan float32) {
	var wg sync.WaitGroup
	wg.Add(len(in))

	name := []string{"noun", "verb"}
	for i := 0; i < len(in); i++ {
		go func(n int, c <-chan float32) {
			for avg := range c {
				fmt.Printf("Average number of %ss per input text: %f\n", name[n], avg)
			}
			wg.Done()
		}(i, in[i])
	}
	wg.Wait()
}
