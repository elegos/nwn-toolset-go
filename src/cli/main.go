package main

import (
	"aurora/file/erf"
	"cli/command"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintf(os.Stdout, "%s [flags] moduleFileName\n", os.Args[0])
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprint(os.Stdout, "Flags:\n")
		flag.PrintDefaults()
		fmt.Fprint(os.Stdout, "\n")
	}

	extract := flag.String("extract", "", "ACTION: extract the file contents in the specified directory")
	readGff := flag.Bool("read-gff", false, "ACTION: explain the GFF files found in the module")
	flag.Parse()

	if flag.NArg() < 1 {
		flag.Usage()

		return
	}

	fileName := flag.Arg(0)

	stat, err := os.Stat(fileName)

	if err != nil || stat == nil || stat.IsDir() {
		fmt.Println(fmt.Sprintf("%s: directory or not existing file", fileName))
		fmt.Println("")
		flag.Usage()

		return
	}

	// Some stats
	module, err := erf.FromFile(fileName)
	if err != nil {
		panic(err)
	}

	command.DescribeErf(&module)

	if *extract != "" {
		command.ExtractErf(&module, *extract)
	}

	if *readGff {
		fmt.Println("")
		fmt.Println("===================================")
		fmt.Println("GFF files found in module")
		fmt.Println("")

		command.ReadGffFromErf(&module)
	}
}
