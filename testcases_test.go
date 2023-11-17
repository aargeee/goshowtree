package goshowtree_test

import (
	"testing/fstest"

	"github.com/Rahul-NITD/goshowtree"
)

var fs_one_file = fstest.MapFS{
	"mydata": {Data: []byte("HI")},
}

var fs_two_files = fstest.MapFS{
	"file1": {Data: []byte("")},
	"file2": {Data: []byte("")},
}

var fs_one_dir = fstest.MapFS{
	"dir1/file2": {Data: []byte{}},
}

var fs_multiple_files = fstest.MapFS{
	"file0":      {Data: []byte{}},
	"dir1/file1": {Data: []byte{}},
	"dir1/file2": {Data: []byte{}},
	"dir2/file3": {Data: []byte{}},
	"dir2/file4": {Data: []byte{}},
}

var tree_one_file = goshowtree.Tree{
	Name: ".",
	SubTree: []goshowtree.Tree{
		{
			Name:    "mydata",
			SubTree: []goshowtree.Tree{},
			IsDir:   false,
		},
	},
	IsDir: true,
}

var tree_two_files = goshowtree.Tree{
	Name: ".",
	SubTree: []goshowtree.Tree{
		{
			Name:    "file1",
			SubTree: []goshowtree.Tree{},
			IsDir:   false,
		},
		{
			Name:    "file2",
			SubTree: []goshowtree.Tree{},
			IsDir:   false,
		},
	},
	IsDir: true,
}

var tree_one_dir = goshowtree.Tree{
	Name: ".",
	SubTree: []goshowtree.Tree{
		{
			Name: "dir1",
			SubTree: []goshowtree.Tree{
				{
					Name:    "file2",
					SubTree: []goshowtree.Tree{},
					IsDir:   false,
				},
			},
			IsDir: true,
		},
	},
	IsDir: true,
}

var tree_multiple_files = goshowtree.Tree{
	Name: ".",
	SubTree: []goshowtree.Tree{
		{
			Name: "dir1",
			SubTree: []goshowtree.Tree{
				{
					Name:    "file1",
					SubTree: []goshowtree.Tree{},
					IsDir:   false,
				},
				{
					Name:    "file2",
					SubTree: []goshowtree.Tree{},
					IsDir:   false,
				},
			},
			IsDir: true,
		},
		{
			Name: "dir2",
			SubTree: []goshowtree.Tree{
				{
					Name:    "file3",
					SubTree: []goshowtree.Tree{},
					IsDir:   false,
				},
				{
					Name:    "file4",
					SubTree: []goshowtree.Tree{},
					IsDir:   false,
				},
			},
			IsDir: true,
		},
		{
			Name:    "file0",
			SubTree: []goshowtree.Tree{},
			IsDir:   false,
		},
	},
	IsDir: true,
}
