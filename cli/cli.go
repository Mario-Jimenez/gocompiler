package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	// VM flag value
	VM string
	// InstructionsCode flag value
	InstructionsCode string
)

// Parse cli
func Parse() {
	codepath := filepath.Join("code", "monkeycode.txt")

	flag.StringVar(&VM, "vm", "", "Path to Monkey Virtual Machine binary")
	flag.StringVar(&InstructionsCode, "code", codepath, "Path to Monkey Instructions Code file")
	flag.Parse()

	if len(VM) == 0 {
		fmt.Println("Missing -vm flag")
		flag.Usage()
		os.Exit(0)
	}

	if _, err := os.Stat(VM); err != nil {
		fmt.Println(err)
		flag.Usage()
		os.Exit(0)
	}
}
