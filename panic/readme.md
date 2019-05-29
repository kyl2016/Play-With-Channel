# panic

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

