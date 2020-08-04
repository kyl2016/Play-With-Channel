package sample

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	s := struct {
		Field1 string
		Field2 int32
		m      map[string]interface{}
	}{
		Field1: "abc",
		Field2: 123,
		m:      map[string]interface{}{"F3": "def", "F4": 1.32},
	}

	buf, _ := json.Marshal(s)
	fmt.Printf("%s", buf)
	// 没有打印 m
	// {"Field1":"abc","Field2":123}
}

type pair struct {
	Key string
	Val interface{}
}

func TestMarshal2(t *testing.T) {
	s := struct {
		Field1 string
		Field2 int32
		a      []pair
	}{
		Field1: "abc",
		Field2: 123,
		a:      []pair{{"F3", "def"}, {"F4", 1.23}},
	}

	buf, _ := json.Marshal(s)
	fmt.Printf("%s", buf)
	// 没有打印 a
	// {"Field1":"abc","Field2":123}
}

func TestMarshal3(t *testing.T) {
	s := struct {
		m map[string]interface{}
	}{
		map[string]interface{}{"F1": "123", "F2": 1.1},
	}

	buf, _ := json.Marshal(s)
	fmt.Printf("%s", buf)
	// {}
}

func TestMarshal4(t *testing.T) {
	m := map[string]interface{}{"F1": "123", "F2": 1.1}

	buf, _ := json.Marshal(m)
	fmt.Printf("%s", buf)
	// {"F1":"123","F2":1.1}
}

func TestConcat(t *testing.T) {
	s := struct {
		Field1 string
		Field2 int32
		m      map[string]interface{}
	}{
		Field1: "abc",
		Field2: 123,
		m:      map[string]interface{}{"F3": "def", "F4": 1.32},
	}

	buf, _ := json.Marshal(s)
	fmt.Printf("%s\n", buf)
	dst := bytes.NewBuffer(buf)
	src, _ := json.Marshal(s.m)

	json.Compact(dst, src)
	fmt.Printf("%s\n", dst.Bytes())

	// 没有打印 m
	// {"Field1":"abc","Field2":123}{"F3":"def","F4":1.32}
}

func TestCombine(t *testing.T) {
	s := struct {
		Field1 string
		Field2 int32
		m      map[string]interface{}
	}{
		Field1: "abc",
		Field2: 123,
		m:      map[string]interface{}{"F3": "def", "F4": 1.32},
	}

	m2 := s.m
	m2["Field1"] = s.Field1
	m2["Field2"] = s.Field2

	buf, _ := json.Marshal(m2)
	fmt.Printf("%s\n", buf)
	// {"F3":"def","F4":1.32,"Field1":"abc","Field2":123}

	s.m["F3"] = "defg"
	buf, _ = json.Marshal(m2)
	fmt.Printf("%s\n", buf)
	// {"F3":"defg","F4":1.32,"Field1":"abc","Field2":123}
}
