# trace

## 从 stack trace 中获取 function、line、file 信息

[Build a Go package that annotates errors with stack traces. (18 November 2019)](https://www.komu.engineer/blogs/golang-stacktrace/golang-stacktrace)

[sample](stacktrace_from_runtime/main1.go)

Recovering from panics

```cassandraql
defer func(){
  if err := recover(); err != nil {
    fmt.Println(errors.Wrap(err, 2).ErrorStack())
  }
}
```

### go-errors/errors

[go-errors blog](https://www.bugsnag.com/blog/go-errors)

