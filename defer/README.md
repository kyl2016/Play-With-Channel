## defer
 
defer 每次执行时，Go 语言会把它携带的 defer 函数及其参数值另行存储到一个栈（FILO），延迟到函数 return 前执行。

### return 会做几件事：
1. 给返回值赋值
2. 调用 defer 表达式
3. 返回给调用函数 （此时是原子操作）

```$xslt
func main() {
	r := test()
	println(r)
}

func test() (i int) {
	defer func() {
		i = i * 2
		println(i)
	}()

	println(i)

	return 3
}
```
return 3 执行过程：
1. 给返回值 i 赋值，即 i=3
2. 调用 defer 表达式，由于在闭包中引用了 i 变量，因此当执行 `i=i * 2`时 i 已经变为 3，defer 执行完后，i 变为6。
3. 返回给调用函数，此时 i 已经为6。
因此执行结果为：
```$xslt
0
6
6
```

### 关于 defer 的几个例子

- [示例一](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/filo/main1.go)：
```$
    for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in  for [%d]\n", i)
	}
```
defer 调用了函数 fmt.Printf，其中 i 作为参数传入，会进行值拷贝，因此输出：
```$xslt
defer in  for [2]
defer in  for [1]
defer in  for [0]
```

- [示例二](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/filo/main2.go)：
```$xslt
    for i := 0; i < 3; i++ {
		defer func() {
			fmt.Printf("defer in  for [%d]\n", i)
		}()
	}
```
defer 调用无参函数 func，且引用了外部变量 i，因此是闭包，那么会将 i 的地址传进函数。循环执行完毕时，i 已经变为 3，因此输出为：
```$xslt
defer in  for [3]
defer in  for [3]
defer in  for [3]
```

- [示例三](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/filo/main3.go)：

```$xslt
    for i := 0; i < 3; i++ {
		defer func(i2 int) {
			fmt.Printf("defer in  for [%d]\n", i2)
		}(i)
	}
```
将 i 作为参数传入函数，因此也会进行值拷贝，结果与示例一相同。

- [示例四](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/filo/main4.go)：
```$xslt
	for i := 0; i < 3; i++ {
		go func() {
			defer fmt.Printf("defer in  for [%d]\n", i)
		}()
	}

	time.Sleep(time.Millisecond * 100)
```
由于 sleep 了一会，与示例二结果一样。

- [示例五](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/return2.go)：
```$xslt
func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
```
defer 执行的闭包只是改变了变量 t 的值，而在此之前，已经将 t 的值赋值给 r，之后再修改 t 的值与 r 就没关系了。

- [示例五](https://github.com/kyl2016/Play-With-Golang/tree/master/defer/callFunc.go)：
```$xslt
func main() {
	var d *People
	defer deferExec(d.Println)
	d = new(People)
}

type People struct {
	Age uint32
}

func (t *People) Println() {
	println(t.Age)
}

func deferExec(f func()) {
	f()
}
```
`defer deferExec(d.Println)` 中的`d`已经求值了，尽管后面调用了`d = new(People)`，也不会改变。下面的例子更明确一些：
```
func main() {
	var d *Data
	defer deferExec2(d.S.F)
	d = new(Data)
}

type Sub struct {
	F func()
}

type Data struct {
	Max int
	S   Sub
}

func (t *Data) Println() {
	println(t.Max)
}

func deferExec2(f func()) {
	f()
}
```
执行到`defer deferExec2(d.S.F)`就引发 panic 了，因为要求值 `d.S`，猜猜下面输出什么？
```$xslt
func main() {
	var d *Data3
	d = new(Data3)
	d.Max = 1
	println("Max:", d.Max)

	defer deferExec3(d.Println)

	d = new(Data3)
	d.Max = 2
	println("Max:", d.Max)
}

type Data3 struct {
	Max int
}

func (t *Data3) Println() {
	println("Println:", t.Max)
}

func deferExec3(f func()) {
	f()
}
```

- 参考：

郝林老师的专栏：[Go语言核心36讲](https://time.geekbang.org/column/article/40889)

### [defer-panic-and-recover](https://blog.golang.org/defer-panic-and-recover)

1. A deferred function's arguments are evaluated when the defer statement is evaluated.

In this example, the expression "i" is evaluated when the Println call is deferred. The deferred call will print "0" after the function returns.
```cassandraql
func a() {
    i := 0
    defer fmt.Println(i)
    i++
    return
}
```

2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.

This function prints "3210":
```cassandraql
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Print(i)
    }
}
```

3. Deferred functions may read and assign to the returning function's named return values.
In this example, a deferred function increments the return value i after the surrounding function returns. Thus, this function returns 2:
```
func c() (i int) {
    defer func() { i++ }()
    return 1
}
```

