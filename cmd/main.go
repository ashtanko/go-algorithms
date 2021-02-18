package main

import (
	"flag"
	"fmt"
	"github.com/ashtanko/go-algorithms/internal/lisp"
	"os"
)

var (
	printSExpr = flag.Bool("sexpr", false, "always print S-expressions")
	stackDepth = flag.Int("depth", 1e5, "maximum call depth; 0 means no limit")
)

var loading bool

func main() {
	flag.Parse()
	lisp.Config(*printSExpr)
	context := lisp.NewContext(*stackDepth)
	loading = true
	for _, file := range flag.Args() {
		load(context, file)
	}
}

// load reads the named source file and parses it within the context.
func load(context *lisp.Context, file string) {
	fd, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer fd.Close()
	//parser := lisp.NewParser(bufio.NewReader(fd))
	//input(context, parser, "")
}
