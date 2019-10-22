package main

type User struct {
	Id   int
	Name string
	Age  int
}

func (u *User) Set(name string, age int) string {
	u.Name = name
	u.Age = age

	return u.Name
}
