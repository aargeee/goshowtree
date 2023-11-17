package goshowtree

import "io/fs"

type Tree struct {
	Name    string
	SubTree []Tree
	IsDir   bool
}

func BuildTree(fsys fs.FS) Tree {
	return Tree{
		Name:    ".",
		SubTree: parseFolder(fsys),
		IsDir:   true,
	}
}

func parseFolder(fsys fs.FS) []Tree {
	var tr []Tree
	names, _ := fs.Glob(fsys, "*")
	for _, name := range names {
		fi, _ := fs.Stat(fsys, name)
		sbfs, _ := fs.Sub(fsys, name)

		tr = append(tr, Tree{
			Name:    name,
			IsDir:   fi.IsDir(),
			SubTree: parseFolder(sbfs),
		},
		)
	}
	return tr
}
