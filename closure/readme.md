# closure

- Transfer the reference of variable i. [sample](../defer/closure.go)
```$
i := 0
defer func() {
    fmt.Println(i)
}()
```

- Share the variable x in the same object a. [sample](shareVariable.go)
```$xslt
var a = Accumulator()
fmt.Printf("%d\n", a(1))
fmt.Printf("%d\n", a(10))
fmt.Printf("%d\n", a(100))

func Accumulator() func(int) int {
	var x int
	return func(delta int) int {
		fmt.Printf("(%+v, %+v) - ", &x, x)
		x += delta
		return x
	}
}

```