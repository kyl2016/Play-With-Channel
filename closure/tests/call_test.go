package tests

import (
	"fmt"
	"testing"
	"time"
)

type Stage struct {
	Name string
}

func (s *Stage) Exec() {
	time.Sleep(time.Second)
	fmt.Println(s.Name)
}

func (s *Stage) AsynExec() {
	go func() {
		s.Exec()
	}()
}

func (s *Stage) AsynExec2() {
	name := s.Name
	go func() {
		time.Sleep(time.Second)
		fmt.Println("AsynExec2", name)
	}()
}

func (s *Stage) AsynExec3() {
	go func(tmp *Stage) {
		tmp.Exec()
	}(s)
}

func TestCall(t *testing.T) {
	s1 := Stage{"1"}
	s2 := Stage{"2"}

	ss := []Stage{s1, s2}

	// for _, s := range ss {
	// 	s.Exec()
	// }

	// for _, s := range ss {
	// 	s.AsynExec()
	// }

	// for _, s := range ss {
	// 	s.AsynExec2()
	// }

	// var j Stage
	// for i := range ss {
	// 	j = ss[i]
	// 	fmt.Printf("%p\n", &j)
	// 	j.AsynExec3()
	// }

	// for i := range ss {
	// 	fmt.Printf("%p\n", &ss[i])
	// 	ss[i].AsynExec3()
	// }

	// for _, s := range ss {
	// 	printS(s)
	// }

	for _, s := range ss {
		printSA(&s)
	}

	time.Sleep(time.Second * 1)
}

func printS(s Stage) {
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println(s.Name)
	}()
}

func printSA(s *Stage) {
	go func() {
		time.Sleep(time.Millisecond)
		fmt.Println(s.Name)
	}()
}
