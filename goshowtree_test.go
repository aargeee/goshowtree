package goshowtree_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	"github.com/Rahul-NITD/goshowtree"
)

func TestGoShowTree(t *testing.T) {
	t.Run("fs having one file", func(t *testing.T) {
		fs := fstest.MapFS{
			"mydata": {Data: []byte("HI")},
		}
		got := goshowtree.BuildTree(fs)
		want := goshowtree.Tree{
			Name: ".",
			SubTree: []goshowtree.Tree{
				{
					Name:    "mydata",
					SubTree: []goshowtree.Tree{},
				},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v != %+v", got, want)
		}
	})
	t.Run("fs having one file with different name", func(t *testing.T) {
		fs := fstest.MapFS{
			"somedifferentName": {Data: []byte("")},
		}
		got := goshowtree.BuildTree(fs)
		want := goshowtree.Tree{
			Name: ".",
			SubTree: []goshowtree.Tree{
				{
					Name:    "somedifferentName",
					SubTree: []goshowtree.Tree{},
				},
			},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v != %+v", got, want)
		}
	})
}
