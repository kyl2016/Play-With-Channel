package main

func main() {
	print(nil)
}

func print(s *myStruct) {
	if s != nil {
		println(s.Name())
	}

	// ...

	if s != nil {
		println(s.Age())
	}
}

type myStruct struct {
	name string
	age  uint8
}

func (s *myStruct) Name() string {
	return s.name
}

func (s *myStruct) Age() uint8 {
	return s.age
}
