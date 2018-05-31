package main

import (
	"os"
)

func main() {
	cli := &CLI{
		StdOut: os.Stdout,
		StdErr: os.Stderr,
	}
	os.Exit(cli.Run(os.Args))
}
