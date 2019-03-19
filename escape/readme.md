# escape

go build -gcflags '-m -l' main.go

```$xslt
./main.go:12:9: &t escapes to heap
./main.go:11:2: moved to heap: t
./main.go:7:14: *x escapes to heap
./main.go:7:13: main ... argument does not escape

```