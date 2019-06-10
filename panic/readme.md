# panic 运行时恐慌

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
  
      
      