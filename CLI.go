package goshowtree

import (
	"fmt"
	"io"
	"strings"
)

const ColorStringRootName = "\u001b[34;1m%s\u001b[0m\n"
const StringRootName = "%s\n"

var Color bool

func ShowTree(w io.Writer, tree Tree, color bool) {
	Color = color
	if Color {
		fmt.Fprintf(w, ColorStringRootName, tree.Name)
	} else {
		fmt.Fprintf(w, StringRootName, tree.Name)
	}
	for _, sbtr := range tree.SubTree {
		rec(w, sbtr, 0)
	}
}

func rec(w io.Writer, tree Tree, depth int) {
	if tree.IsDir {
		if Color {
			fmt.Fprintf(w, "%s├── "+ColorStringRootName, strings.Repeat("    ", depth), tree.Name)
		} else {
			fmt.Fprintf(w, "%s├── "+StringRootName, strings.Repeat("    ", depth), tree.Name)
		}
		for _, sbtr := range tree.SubTree {
			rec(w, sbtr, depth+1)
		}
	} else {
		fmt.Fprintf(w, "%s├── %s\n", strings.Repeat("    ", depth), tree.Name)
	}
}
