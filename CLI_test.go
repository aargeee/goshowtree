package goshowtree_test

import (
	"bytes"
	"testing"

	"github.com/Rahul-NITD/goshowtree"
)

func TestCLI(t *testing.T) {
	t.Run("single file print", func(t *testing.T) {
		out := &bytes.Buffer{}
		tree := tree_one_file
		goshowtree.ShowTree(out, tree, false)
		want :=
			`.
├── mydata
`
		got := out.String()
		if want != got {
			t.Errorf("got \n%s\nwanted\n%s", got, want)
		}
	})
	t.Run("two file print", func(t *testing.T) {
		out := &bytes.Buffer{}
		tree := tree_two_files
		goshowtree.ShowTree(out, tree, false)
		want :=
			`.
├── file1
├── file2
`
		got := out.String()
		if want != got {
			t.Errorf("got \n%s\nwanted\n%s", got, want)
		}
	})
	t.Run("one folder print", func(t *testing.T) {
		out := &bytes.Buffer{}
		tree := tree_one_dir
		goshowtree.ShowTree(out, tree, false)
		want :=
			`.
├── dir1
    ├── file2
`
		got := out.String()
		if want != got {
			t.Errorf("got \n%q\nwanted\n%q", got, want)
		}
	})
	t.Run("multiple files and folders print", func(t *testing.T) {
		out := &bytes.Buffer{}
		tree := tree_multiple_files
		goshowtree.ShowTree(out, tree, false)
		want :=
			`.
├── dir1
    ├── file1
    ├── file2
├── dir2
    ├── file3
    ├── file4
├── file0
`
		got := out.String()
		if want != got {
			t.Errorf("got \n%s\nwanted\n%s", got, want)
		}
	})
}
