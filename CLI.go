package goshowtree

import (
	"fmt"
	"io"
	"strings"
)

func ShowTree(w io.Writer, tree Tree) {
	fmt.Fprintf(w, "%s\n", tree.Name)
	for _, sbtr := range tree.SubTree {
		rec(w, sbtr, 0)
	}
}

func rec(w io.Writer, tree Tree, depth int) {
	if tree.IsDir {
		fmt.Fprintf(w, "%s|-- %s\n", strings.Repeat("|   ", depth), tree.Name)
		for _, sbtr := range tree.SubTree {
			rec(w, sbtr, depth+1)
		}
	} else {
		fmt.Fprintf(w, "%s|-- %s\n", strings.Repeat("|   ", depth), tree.Name)
	}
}
