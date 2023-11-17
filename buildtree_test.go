package goshowtree_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Rahul-NITD/goshowtree"
)

func TestGoShowTree(t *testing.T) {

	cases := []struct {
		testName string
		testFS   fstest.MapFS
		want     goshowtree.Tree
	}{
		{
			testName: "fs having one file",
			testFS: fstest.MapFS{
				"mydata": {Data: []byte("HI")},
			},
			want: goshowtree.Tree{
				Name: ".",
				SubTree: []goshowtree.Tree{
					{
						Name:    "mydata",
						SubTree: []goshowtree.Tree{},
						IsDir:   false,
					},
				},
				IsDir: true,
			},
		},
		{
			testName: "fs having one file with different name",
			testFS: fstest.MapFS{
				"somedifferentName": {Data: []byte("")},
			},
			want: goshowtree.Tree{
				Name: ".",
				SubTree: []goshowtree.Tree{
					{
						Name:    "somedifferentName",
						SubTree: []goshowtree.Tree{},
						IsDir:   false,
					},
				},
				IsDir: true,
			},
		},
		{
			testName: "fs having two files",
			testFS: fstest.MapFS{
				"file1": {Data: []byte("")},
				"file2": {Data: []byte("")},
			},
			want: goshowtree.Tree{
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
			},
		},
		{
			testName: "fs having one directory with one file",
			testFS: fstest.MapFS{
				"dir1/file2": {Data: []byte{}},
			},
			want: goshowtree.Tree{
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
			},
		},
		{
			testName: "fs having two directory with two file each",
			testFS: fstest.MapFS{
				"file0":      {Data: []byte{}},
				"dir1/file1": {Data: []byte{}},
				"dir1/file2": {Data: []byte{}},
				"dir2/file3": {Data: []byte{}},
				"dir2/file4": {Data: []byte{}},
			},
			want: goshowtree.Tree{
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
			},
		},
	}

	for _, test := range cases {
		t.Run(test.testName, func(t *testing.T) {
			got := goshowtree.BuildTree(test.testFS)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %+v != %+v", got, test.want)
			}
		})
	}

}
