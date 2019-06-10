## defer
 
defer 每次执行时，Go 语言会把它携带的 defer 函数及其参数值另行存储到一个栈（FILO）。

[示例一](./filo/main1.go)：
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

[示例二](./filo/main2.go)：
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

[示例三](./filo/main3.go)：

```$xslt
    for i := 0; i < 3; i++ {
		defer func(i2 int) {
			fmt.Printf("defer in  for [%d]\n", i2)
		}(i)
	}
```
将 i 作为参数传入函数，因此也会进行值拷贝，结果与示例一相同。

[示例四](./filo/main4.go)：
```$xslt
	for i := 0; i < 3; i++ {
		go func() {
			defer fmt.Printf("defer in  for [%d]\n", i)
		}()
	}

	time.Sleep(time.Millisecond * 100)
```
由于 sleep 了一会，与示例二结果一样。