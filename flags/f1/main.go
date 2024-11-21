package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "0.0.1"

//var Usage = func() {
//	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n Version\n", os.Args[0], version)
//	//fmt.Fprintf(flag.CommandLine.Output(), "", )
//	//flag.PrintDefaults()
//}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n Version %s\n", os.Args[0], version)
		flag.PrintDefaults()
	}

	flag.Parse()
}
