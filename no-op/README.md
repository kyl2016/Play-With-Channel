# no-op


Should check s is nil when every call its method in [sample](./main1.go)
```$xslt

func print(s *myStruct) {
	if s != nil {
		println(s.Name())
	}

	// ...

	if s != nil {
		println(s.Age())
	}
}

```

Or check only in its method internal, [sample](./main2.go)
```$xslt

func (s *myStruct2) Name() string {
	if s == nil {
		return ""
	}

	return s.name
}
```
then, needn't check when calling
```$xslt
func print2(s *myStruct2) {
	println(s.Name())
}

```