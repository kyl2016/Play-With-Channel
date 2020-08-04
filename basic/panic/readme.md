# panic (运行时)恐慌



- recover is ONLY valid in current goroutine. 
```
go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success. err:", err)
			}
		}()
// ...
}()
```

One application of recover is to shut down a failing goroutine inside a server without killing the other executing goroutines.
In this [example](oneGoroutine.go) , if do(work) panics, the result will be logged and the goroutine will exit cleanly without disturbing the others.


- 从 panic 被引发到程序终止运行的大致过程是什么？
    > 某个函数中的某行代码有意或无意地引发一个 panic。
    初始的 panic 详情会被建立起来，
    
      
## 每个 goroutine 需要单独调用 recover
 主函数的 recover 无法捕获其它 goroutines 的 panic，因此每个 goroutine 单独处理，参见 [multiGoroutines](multiGoroutines.go)
 
## panic-recover 与 try-catch 的区别
这是两种完全不同的异常处理机制。
Go 语言的异常处理机制是两层的，defer 和 recover 可以处理**意外**的异常，而 error 接口及相关体系处理**可预期**的异常。
 
 
## 参考

[Golang: 深入理解panic and recover](https://ieevee.com/tech/2017/11/23/go-panic.html)