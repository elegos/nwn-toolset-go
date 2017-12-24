package main

import (
	"aurora/file/erf"
	"cli/command"
	"flag"
)

func extract(module *erf.ERF) {

}

func main() {
	fileName := flag.String("file", "", "the module's file name")
	extractFlag := flag.String("extract", "", "ACTION: extract the file contents in the specified directory")

	flag.Parse()

	// Some stats
	module := erf.FromFile(*fileName)

	command.DescribeErf(&module)

	if *extractFlag != "" {
		command.ExtractErf(&module, *extractFlag)
	}
}
