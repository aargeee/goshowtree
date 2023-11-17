package main

import (
	"fmt"
	"io/fs"
	"testing/fstest"
)

func main() {
	tempFS := fstest.MapFS{
		// "tmp":           {Mode: fs.ModeDir},
		"tmp/something": {Data: []byte{}},
	}
	fi, _ := fs.Stat(tempFS, "tmp")
	sub, _ := fs.Sub(tempFS, "tmp")
	sb, _ := fs.Glob(sub, "*")
	fmt.Println(fi.IsDir())
	fmt.Println(sb)

}
