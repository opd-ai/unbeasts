package main

import (
	"fmt"

	"github.com/opd-ai/untemplate"
)

func run() string {
	return untemplate.Summary()
}

func main() {
	fmt.Println(run())
}
