# stack and heap

## Which stays on the stack?

sample 1:
```
func main() {
	b := read()
	println(len(b))
}

func read() []byte {
	r := make([]byte, 10)
	return r
}
```

sample 2:
```cassandraql
func main() {
	s := make([]byte, 10)
	write(s)
	println(len(s))
}

func write(s []byte) {
	s[0] = 'a'
}
```

Sample 1 is bad, s is on the heap, this will frequently be going to the heap.
Sample 2 is good, s 在栈上。

## io.Reader

type Reader interface {
    Read(p []byte) (n int, err error)
}

Instead of 

type Reader interface {
    Read(n int) (b []byte, err error)
}

You make the slice and pass the slice into the read method and it returns a number to tell you how much of your slice it filled.
The second example, we'll have much garbage on the heap.

