# plugin

先编译 pkg 到 pkg.so，然后再运行 main.go

// 注意编译 pkg 和 main.go 的参数一定要一致，如果后者有`-race -gcflags "all=-N -l"`，则 pkg 的编译也要有

[Package plugin](https://golang.org/pkg/plugin/)