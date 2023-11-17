package goshowtree_test

import (
	"fmt"
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
			testFS:   fs_one_file,
			want:     tree_one_file,
		},
		{
			testName: "fs having two files",
			testFS:   fs_two_files,
			want:     tree_two_files,
		},
		{
			testName: "fs having one directory with one file",
			testFS:   fs_one_dir,
			want:     tree_one_dir,
		},
		{
			testName: "fs having two directory with two file each",
			testFS:   fs_multiple_files,
			want:     tree_multiple_files,
		},
	}

	for _, test := range cases {
		t.Run(test.testName, func(t *testing.T) {
			tree := goshowtree.BuildTree(test.testFS)
			// Model based testing
			got := fmt.Sprintf("%+v", tree)
			want := fmt.Sprintf("%+v", test.want)
			if got != want {
				t.Errorf("got\n%s\nwant\n%s\n", got, want)
			}
		})
	}

}
