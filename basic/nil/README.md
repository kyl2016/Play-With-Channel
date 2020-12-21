# nil

## nil is a predeclared identifier in go.

## cannot set 'nil' to string

The nil constant has **no type** — so it cannot substitute for a string.

    package main
    
    func main() {
        temp := []string{}
        temp = append(temp, nil)
    }
    
    Output
    
    # command-line-arguments
    C:\programs\file.go:8: cannot use nil as type string in append

## nil can represent zero values of many types

- pointer types (including type-unsafe ones)
- map types
- slice types
- function types
- channel types
- interface types

## predeclared nil is not a keyword in Go

The predecalred nil can be shadowed.

```
nil := 123
fmt.Println(nil) // 123

// The dollowing line fails to compile,
// for nil represents an int value now 
// in this scope.
var _ map[string]int = nil
```

## the size of nil values with types of different kinds may be different

## two nil values of two different types may be not comparable

```
// Compilation failure reason: mismatched types.
// neither operand can be implicitly converted to the type of the other.
var _ = (*int)(nil) == (*bool)(nil) // error
var _ = (chan int)(nil) == (chan bool)(nil) // error
```

```
type IntPtr *int
// The underlying of type IntPtr is *int.
var _ = IntPtr(nil) == (*int)(nil)

// Every type in Go implements interface{} type.
var _ = (interface{})(nil) == (*int)(nil)

// Values of a directional channel type can be converted to the bidirectional channel type which has the same element type.
var _ = (chan int)(nil) == (chan<- int)(nil)
var _ = (chan int)(nil) == (<-chan int)(nil)
```

## two nil values of the same type may be not comparable

```
// illegal
var _ = ([]int)(nil) == ([]int)(nil)
var _ = (map[string]int)(nil) == (map[string]int)(nil)
var _ = (func())(nil) == (func())(nil)
```

But any types of the above mentioned incomparable types can be compared with
the bare nil identifier.
```
// The following lines compile okay.
var _ = ([]int)(nil) == nil
var _ = (map[string]int)(nil) == nil
var _ = (func())(nil) == nil
```

## ！！！ two nil values may be not equal

If one of the two compared nil values is an interface{} value and the other is not,
assume they are comparable, then the comparison result is always false. The
reson is the not-interface value will be [converted to the type of interface value](https://go101.org/article/interface.html#boxing)
before making the comparison. **The converted interface value has a concrete
dynamic type but the other interface value has not**.

```
fmt.Println( (interface{})(nil) == (*int)(nil) ) // false
```

## Retrieving elements from nil maps will not panic

```
fmt.Println( (map[string]int)(nil)["key"] ) // 0
fmt.Println( (map[int]bool)(nil)[123]) // false
fmt.Println( (map[int]*int64)(nil)[123]) // <nil>
```

## predeclared nil has not a default type

In fact, the predecared nil is the only untyped value who has not a default type in Go. There must be sufficient information for compiler to deduce(推断) the type of a nil from context.

in go, for simplicity and convenience, nil is designed as an identifier which can be used to represent the zero values of some kinds of types. it is **not a single value**. it can represent many values with different memory layouts.

## refer

[go101/nil](https://go101.org/article/nil.html)