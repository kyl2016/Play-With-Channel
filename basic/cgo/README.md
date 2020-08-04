# cgo


## c call go func

go functions can be exported for use by c code in the following way:

```cassandraql
//export myfunction
func myfunction(arg1, arg2 int, arg3 string) int64 {...}

//export myfunction2
func myfunction2(arg1, arg2 int, arg3 string) (int64, *c.char) {...}
```

they will be available in the c code as:

```cassandraql
extern goint64 myfunction(int arg1, int arg2, gostring arg3);
extern struct myfunction2_return myfunction2(int arg1, int arg2, gostring arg3);
```

found in the `_cgo_export.h` generated header.

[sample](c_to_go/c_to_go.go)

## strings and things

strings in c are represented by a zero-terminated array of chars.

conversion:
* C.CString
* C.GoString
* C.GoStringN

**free**
Memory allocations made by C code are not known to Go's memory manager. Should free it by calling C.free.

## callback

[sample](./callback/main.go)
[refer](https://dev.to/mattn/call-go-function-from-c-function-1n3)