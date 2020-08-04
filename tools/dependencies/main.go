package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

// go list -m -json all

func main() {
	depFile := "/Users/kang/go/src/github.com/kyl2016/Play-With-Golang/pkg/etcd/test/etcd.dep.json"
	depFile = "/Users/kang/go/src/github.com/kyl2016/Play-With-Golang/image/crop/dep.json"
	depFile = "/Users/kang/go/src/github.com/kyl2016/Play-With-Golang/pkg/fileNotify/dep.json"

	srcRoot := "/Users/kang/go/pkg/mod/cache/download/"
	saveRoot := "/tmp/download"

	file, err := os.Open(depFile)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	count := 0

	r, err := regexp.Compile("Path\": \"(.*?)\",")

	for _, eachline := range txtlines {
		if strings.Index(eachline, "Path") > 0 {
			count++

			fmt.Println(r.FindStringSubmatch(eachline)[1])
			matches := r.FindStringSubmatch(eachline)
			if len(matches) > 0 {
				cp(srcRoot, r.FindStringSubmatch(eachline)[1], saveRoot)
			} else {
				fmt.Println(r.FindStringSubmatch(eachline)[1])
			}
		}
	}

	println(count)
}

func cp(srcRoot, relativeFile, saveRoot string) {
	source := path.Join(srcRoot, relativeFile)

	target := path.Join(saveRoot, relativeFile)

	os.MkdirAll(target, os.ModePerm)

	_, err := os.Stat(source)
	if err != nil {
		fmt.Println("not found", source)
		return
	}

	err = CopyDirectory(source, target)
	if err != nil {
		panic(err)
	}
}

func CopyDirectory(scrDir, dest string) error {
	entries, err := ioutil.ReadDir(scrDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(scrDir, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		stat, ok := fileInfo.Sys().(*syscall.Stat_t)
		if !ok {
			return fmt.Errorf("failed to get raw syscall.Stat_t data for '%s'", sourcePath)
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err := CreateIfNotExists(destPath, 0755); err != nil {
				return err
			}
			if err := CopyDirectory(sourcePath, destPath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err := CopySymLink(sourcePath, destPath); err != nil {
				return err
			}
		default:
			if err := Copy(sourcePath, destPath); err != nil {
				return err
			}
		}

		if err := os.Lchown(destPath, int(stat.Uid), int(stat.Gid)); err != nil {
			return err
		}

		isSymlink := entry.Mode()&os.ModeSymlink != 0
		if !isSymlink {
			if err := os.Chmod(destPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(srcFile, dstFile string) error {
	out, err := os.Create(dstFile)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(srcFile)
	defer in.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateIfNotExists(dir string, perm os.FileMode) error {
	if Exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func CopySymLink(source, dest string) error {
	link, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(link, dest)
}
