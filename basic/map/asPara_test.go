package main

import (
	"fmt"
	"testing"
)

func TestAdd1(t *testing.T) {
	m := map[int][]int{}

	add(m, 1, 1)
	add(m, 1, 2)

	if len(m[1]) == 2 {
		t.Errorf("map add should failed")
	}
}

func add(m map[int][]int, key, value int) {
	if v, ok := m[key]; ok {
		v = append(v, value)
		fmt.Println("v:", v)
		fmt.Printf("m[%d]: %v\n", key, m[key])
	} else {
		m[key] = []int{value}
	}
}

func TestAdd2(t *testing.T) {
	m := map[int][]int{}

	add2(m, 1, 1)
	add2(m, 1, 2)

	if len(m[1]) != 2 {
		t.Errorf("map add failed")
	}
}

func add2(m map[int][]int, key, value int) {
	if v, ok := m[key]; ok {
		m[key] = append(m[key], value)
		fmt.Println("v:", v)
		fmt.Printf("m[%d]: %v\n", key, m[key])
	} else {
		m[key] = []int{value}
	}
}

func TestAdd3(t *testing.T) {
	m := map[int]*[]int{}

	add3(m, 1, 1)
	add3(m, 1, 2)

	if len(*m[1]) != 2 {
		t.Errorf("map add failed")
	}
}

func add3(m map[int]*[]int, key, value int) {
	if v, ok := m[key]; ok {
		*v = append(*v, value)
		fmt.Println("v:", v)
		fmt.Printf("m[%d]: %v\n", key, m[key])
	} else {
		m[key] = &[]int{value}
	}
}

func TestUpdate1(t *testing.T) {
	m := map[int][]int{}

	add2(m, 1, 1)
	update1(m, 1, 0, 2)

	if m[1][0] != 2 {
		t.Errorf("map update failed")
	}
}

func update1(m map[int][]int, key, index, value int) {
	s := m[key]
	s[index] = value
}
