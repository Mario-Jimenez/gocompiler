package main

import (
	"github.com/Mario-Jimenez/gocompiler/cli"
	"github.com/Mario-Jimenez/gocompiler/router"
)

func main() {
	// parse flags
	cli.Parse()
	// start web service
	router.Run()
}
