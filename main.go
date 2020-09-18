package main

import (
	"fmt"
	"os"

	"github.com/Mario-Jimenez/gocompiler/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input, err := antlr.NewFileStream("program.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	inst := parser.NewCompiler(input)

	fmt.Printf("%+v\n", inst)
}
