# slice

[copy_test](./copy_test.go) 中的 TestSliece2，s2=s 未扩容，但 append(s2, 2) 之后，s 未改变，why？

## slice 作为参数
[demo](./value_pointer/main1.go)

### 1. 传递初始化之后的 slice

1.1 函数修改 slice 里面的元素

```
func write(s []byte) {
	s[0] = 'a'
}
```
修改的是 s 对应数组中的值，生效。

1.2 函数在 slice 尾部追加元素
```
func appendToValue(s []byte) {
	s = append(s, 10)
	fmt.Println("------------ appendToValue", len(s), s)
}
```
在 s 尾部追加元素，会触发扩容，则 s 指向了另一个数组，appendToValue 返回后，原来的 slice 指向的数组并未改变，因此追加无效。

这个[demo](./value_pointer/main2.go)打印了过程中slice对应数组第一个元素的地址，一目了然。

### 2. 传递初始化后的 slice 的引用
```
func appendToPointer(s *[]byte) {
	*s = append(*s, 10)
	fmt.Println("------------ appendToPointer", len(*s), *s)
}
```
由于传递的 s 的引用，直接修改了 s 对应的数组。

## references
[关于 Go 中 Map 类型和 Slice 类型的传递 - alfred_zhong](https://www.cnblogs.com/snowInPluto/p/7477365.html)

slice 默认传递的是元素的值，否则传递地址 [示例](./demo3.go)