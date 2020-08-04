# go build

## go build -tags

```$xslt
$ go build -tags linux
$ ./build
linux tag
```

It will exclude go files with `// +build !linux`.

## 不同平台

windows:
    
    $ GOOS=windows GOARCH=386 go build *.go

    $GOOS     $GOARCH     
    darwin     386
    darwin     amd64
    freebsd     386
    freebsd     amd64
    linux     386
    linux     amd64
    linux     arm     incomplete
    windows     386     incomplete
    
## Context

```
type Context struct {
    GOARCH string // target architecture
    GOOS    string  // target operating system
    GOROOT  string  // Go root
    GOPATH  string  // Go path
    // Dir is the caller's working directory, or the empty string to use the current directory of the running process. 
    // In module mode, this is used to locate the main module.
    // If Dir is non-empty, directories passed to Import and ImportDir must be absolute.
    Dir     string  // Go 1.14 
    CgoEnable   bool    // whether cgo files are included
    UseAllFiles bool    // use files regardless of +build lines, file names
    Compiler    string  // compiler to assume when computing target paths

    BuildTags   []string
    RleaseTags  []string    // The last element in ReleaseTags is assumed to be the current release.

    ...
```


## references

[build package](https://golang.org/pkg/go/build/#hdr-Build_Constraints)