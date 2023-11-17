package main

import (
	"os"

	"github.com/Rahul-NITD/goshowtree"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fileSystem := os.DirFS(dir)
	tree := goshowtree.BuildTree(fileSystem, dir)
	goshowtree.ShowTree(os.Stdout, tree)
}
