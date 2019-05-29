package v0_3

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func Test_NotSetSource(t *testing.T) {
	pipe := NewPipeline("demo", true, func(format string, args ...interface{}) {
		log.Printf(format, args...)
	})

	pipe.NewStage("stage1")

	pipe.Run()

	if pipe.State != Failed {
		t.Fatal("should be failed")
	}
}

func Test_Sample(t *testing.T) {
	pipe := NewPipeline("demo", true, func(format string, args ...interface{}) {
		log.Printf(format, args...)
	})

	sourceCh := make(chan interface{}, 100)
	resultCh := make(chan interface{}, 100)
	sumResult := 0

	stage1 := pipe.NewStage("stage1")
	stage1.SetSource("source1", func(p func(in interface{})) {
		for i := 0; i < 3; i++ {
			p(i)
			time.Sleep(time.Millisecond * 300)
		}
	}, sourceCh)

	stage1.AddStep("step1", 1, func(in interface{}) (out interface{}, err error) {
		i := in.(int)
		time.Sleep(time.Millisecond * 200)

		return i * i, nil
	}, sourceCh, resultCh)

	stage1.SetSink("sumStep", 1, func(in interface{}) (err error) {
		sumResult += in.(int)
		return nil
	}, resultCh)

	pipe.Run()

	fmt.Println("sum:", sumResult)
}

func Test_Concurrent(t *testing.T) {
	pipe := NewPipeline("demo", true, func(format string, args ...interface{}) {
		//log.Printf(format, args...)
	})

	sourceCh := make(chan interface{}, 100)
	resultCh := make(chan interface{}, 100)
	sumResult := 0

	stage1 := pipe.NewStage("stage1")
	stage1.SetSource("source1", func(p func(in interface{})) {
		for i := 0; i < 1000000; i++ {
			p(i)
			//time.Sleep(time.Millisecond * 300)
		}
	}, sourceCh)

	stage1.AddStep("step1", 10, func(in interface{}) (out interface{}, err error) {
		i := in.(int)
		//time.Sleep(time.Millisecond * 200)

		return i * i, nil
	}, sourceCh, resultCh)

	stage1.SetSink("sumStep", 10, func(in interface{}) (err error) {
		sumResult += in.(int)
		return nil
	}, resultCh)

	start := time.Now()
	pipe.Run()
	fmt.Println("elapsed:", time.Now().Sub(start).Seconds())

	fmt.Println("sum:", sumResult)

	if pipe.State != Finished {
		t.Fatal("pipeline state should be ", Finished)
	}

	time.Sleep(time.Second * 3) // wait log print finished
}

func Test_Batch(t *testing.T) {
	pipe := NewPipeline("demo", true, func(format string, args ...interface{}) {
		//log.Printf(format, args...)
	})

	sourceCh := make(chan interface{}, 100)
	resultCh := make(chan interface{}, 100)
	var sumResult float32 = 0.0

	stage1 := pipe.NewStage("stage1")
	stage1.SetSource("source1", func(p func(in interface{})) {
		for i := 0; i < 1000000; i++ {
			f := float32(i) * 0.1234
			p(f)
			//time.Sleep(time.Millisecond * 300)
		}
	}, sourceCh)

	stage1.AddBatchStep("step1", 10, 1, func(in []interface{}) (out []interface{}, err error) {
		var results []interface{}

		for _, item := range in {
			i := item.(float32)
			results = append(results, i*i/0.2124123112)
		}

		return results, nil
	}, sourceCh, resultCh)

	stage1.SetBatchSink("sumStep", 10, 1, func(in []interface{}) (err error) {
		for _, item := range in {
			sumResult += item.(float32)
		}
		return nil
	}, resultCh)

	start := time.Now()
	pipe.Run()
	fmt.Println("elapsed:", time.Now().Sub(start).Seconds())

	fmt.Println("sum:", sumResult)

	if pipe.State != Finished {
		t.Fatal("pipeline state should be ", Finished)
	}

	time.Sleep(time.Second * 3) // wait log print finished
}

func Test_MultiStages(t *testing.T) {
	pipe := NewPipeline("demo", true, func(format string, args ...interface{}) {
		log.Printf(format, args...)
	})

	sourceCh := make(chan interface{}, 100)
	resultCh := make(chan interface{}, 100)
	sumResult := 0

	stage1 := pipe.NewStage("stage1")
	stage1.SetSource("source1", func(p func(in interface{})) {
		for i := 0; i < 3; i++ {
			p(i)
			time.Sleep(time.Millisecond * 300)
		}
	}, sourceCh)

	stage1.AddStep("step1", 1, func(in interface{}) (out interface{}, err error) {
		i := in.(int)
		time.Sleep(time.Millisecond * 200)

		return i * i, nil
	}, sourceCh, resultCh)

	stage1.SetSink("sumStep", 1, func(in interface{}) (err error) {
		sumResult += in.(int)
		return nil
	}, resultCh)

	sourceCh2 := make(chan interface{}, 100)
	resultCh2 := make(chan interface{}, 100)
	stage2 := pipe.NewStage("stage2")
	stage2.SetSource("source2", func(p func(in interface{})) {
		for i := 0; i < 3; i++ {
			p(i)
			time.Sleep(time.Millisecond * 300)
		}
	}, sourceCh2)

	stage2.AddStep("step2", 1, func(in interface{}) (out interface{}, err error) {
		i := in.(int)
		time.Sleep(time.Millisecond * 200)

		return i * i, nil
	}, sourceCh2, resultCh2)

	stage2.SetSink("sumStep2", 1, func(in interface{}) (err error) {
		sumResult += in.(int)
		return nil
	}, resultCh2)

	pipe.Run()

	fmt.Println("sum:", sumResult)
}
