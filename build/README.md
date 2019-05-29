# go build -tags

```$xslt
$ go build -tags linux
$ ./build
linux tag
```

It will exclude go files with `// +build !linux`.