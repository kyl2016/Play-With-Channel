# channel

## Samples

题目1： （[答案](sample2.go)）

```go
package main

import (
	"fmt"
	"time"
)

// 一个栅栏接口。
type Barrier interface {
	Wait()
}

// 创建栅栏对象。
func NewBarrier(n int) Barrier {
	// 这里需要填入代码。
}

// 栅栏的实现类型。
type barrier struct {
	// 这里需要填入代码。
}


func (b *barrier) Wait() {
	// 这里需要填入代码。
}

// 测试代码。
func main() {
	num := 10
	// 创建栅栏值。
	b := NewBarrier(num)
	// 需要达到的效果：
	//   前 9 个 goroutine 调用 Wait 方法时被阻塞；
	//   第 10 个 goroutine 调用 Wait 方法后，所有 goroutine 全部被唤醒。
	for i := 0; i < num; i++ {
		go func(i int) {
			fmt.Printf("Wait[%d]\n", i)
			b.Wait()
			fmt.Printf("Done[%d]\n", i)
		}(i)
	}
	time.Sleep(time.Second * 2)
}
```

题目：

需要答题者在上面的 NewBarrier 函数、barrier 类型和 Wait 方法中填入代码，以达到测试代码（main 函数）中所述的效果。

有 4 个约束：

1. 不能使用任何同步工具。
2. 不能使用任何原子操作。
3. 不能使用 context。
4. 可以使用通道（channel），但仅限于非缓冲通道。

题目2：（[答案](sample3.go)）
使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 1A2B3C4D5E6F... 。

题目3：（[答案](sample4.go)）
使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字， 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。

## Remember  all transfer of value on the go channels happends with the copy of value.
[copy value](./samples/copyValue/main.go)

    	c <- g // 将 g 的值（user 的地址）复制到 c
    	g = &user{name: "Ankur Anand", age: 100} // g的地址改变了，c 里面存储的是 g 之前的地址，g 再重新赋值与 c 无关了	
    	
如果，没有修改 g 的值（user 的地址），而是修改 user 的内容，由于 c 中的元素指向的相同的地址，因此内容也会变化

        c := make(chan *user, 5)
        c <- g
        // modify g
        modifyUser(g)
参考：[修改属性值](./samples/editContent/editDirectly.go)