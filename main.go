package main

import (
	"fmt"
	"github.com/claude-betz/mcnaughton-yamada-thompson/parser"
)

var (
	regex = "(a|b)*"
)

func main() {
	fmt.Printf("regex: %s\n", regex)

	ast, err := parser.Parse(&regex)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	fmt.Printf("printing ast\n")
	ast.PrintNFA()
}
