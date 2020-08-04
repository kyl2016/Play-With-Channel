package main

import (
	"fmt"
	"go/build"
)

func main() {
	var toolDir = build.ToolDir
	fmt.Println(toolDir)
	// /usr/local/go/pkg/tool/darwin_amd64

	// panic: architecture letter no longer used
	//	r, err := build.ArchChar("386")
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(r)

	fmt.Println(build.IsLocalImport("."))
	// true

	fmt.Printf("default: %+v\n", build.Default)
	//{GOARCH:amd64 GOOS:darwin GOROOT:/usr/local/go GOPATH:/Users/kang/go Dir: CgoEnabled:true UseAllFiles:false Compiler:gc BuildTags:[] ReleaseTags:[go1.1 go1.2 go1.3 go1.4 go1.5 go1.6 go1.7 go1.8 go1.9 go1.10 go1.11 go1.12 go1.13 go1.14] InstallSuffix: JoinPath:<nil> SplitPathList:<nil> IsAbsPath:<nil> IsDir:<nil> HasSubdir:<nil> ReadDir:<nil> OpenFile:<nil>}
}
