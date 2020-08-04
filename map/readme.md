# map

"comma ok" idiom
```$xslt
seconds, ok := timeZone[tz]
```

## map as parameter

参见[asPara_test](./asPara_test.go)

典型错误如 TestAdd1，调用的 add 方法：
```
func add(m map[int][]int, key, value int) {
	if v, ok := m[key]; ok {
		v = append(v, value)
		fmt.Println("v:", v)
		fmt.Printf("m[%d]: %v\n", key, m[key])
	} else {
		m[key] = []int{value}
	}
}
```
其中，`v = append(v, value)`只是修改了 v 的值，`v, ok := m[key]`，`v` 复制了`m[key]`的内容，
