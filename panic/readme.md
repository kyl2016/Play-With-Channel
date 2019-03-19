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