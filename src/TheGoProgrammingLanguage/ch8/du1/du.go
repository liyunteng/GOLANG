// -*- compile-command: "go run /home/lyt/git/golang/src/TheGoProgrammingLanguage/ch8/du1/du.go "; -*-
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func walkDir(dir string, fileSize chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			// fmt.Println(subdir)
			walkDir(subdir, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}

	return entries
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Println(roots)

	size := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, size)
		}
		close(size)
	}()

	var nfiles, nbytes int64
	for x := range size {
		nfiles++
		nbytes += x
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %d KB\n", nfiles, nbytes/1e3)
}
