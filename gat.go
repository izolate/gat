package main

import (
	"fmt"
	"io"
)

const (
	ExitCodeSuccess int = 0
)

type CLI struct {
	StdOut, StdErr io.Writer
}

func (cli *CLI) Run(args []string) int {
	fmt.Fprintf(cli.StdOut, "Good things come to those who wait")
	return ExitCodeSuccess
}
