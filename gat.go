package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeSuccess int = 0

	AllHelp         = "equivalent to -vET"
	BlankHelp       = "number nonempty output lines, overrides -n"
	EndsHelp        = "display $ at end of each line"
	NumberHelp      = "number all output lines"
	SqueezeHelp     = "Suppress repeated empty output lines"
	TabsHelp        = "display TAB characters as ^I"
	NonprintingHelp = "use ^ and M- notation, except for LFD and TAB"
	UnbufferedHelp  = "(ignored)"
)

var (
	all         = flag.Bool("show-all", false, AllHelp)
	blank       = flag.Bool("number-nonblank", false, BlankHelp)
	ends        = flag.Bool("show-ends", false, EndsHelp)
	number      = flag.Bool("number", false, NumberHelp)
	squeeze     = flag.Bool("squeeze-blank", false, SqueezeHelp)
	tabs        = flag.Bool("show-tabs", false, TabsHelp)
	nonprinting = flag.Bool("show-nonprinting", false, NonprintingHelp)
	unbuffered  = flag.Bool("u", false, UnbufferedHelp)
)

type CLI struct {
	StdOut, StdErr io.Writer
}

func (cli *CLI) Run(args []string) int {
	/* Parse CLI arguments */
	flag.Parse()

	fmt.Fprintf(cli.StdOut, "Good things come to those who wait")
	return ExitCodeSuccess
}
