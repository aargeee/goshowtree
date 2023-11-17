package goshowtree

import "io/fs"

type Tree struct {
	Name    string
	SubTree []Tree
}

func BuildTree(fsys fs.FS) Tree {
	s, _ := fs.Glob(fsys, "*")

	tree := Tree{
		Name:    ".",
		SubTree: []Tree{},
	}

	for _, name := range s {
		tree.SubTree = append(tree.SubTree, Tree{
			Name:    name,
			SubTree: []Tree{},
		})
	}

	return tree
}
