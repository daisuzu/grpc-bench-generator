package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

var (
	path = flag.String("path", ".", "path to the target protocol buffer package")
)

func main() {
	flag.Parse()

	abspath, err := filepath.Abs(*path)
	if err != nil {
		log.Fatal(err)
	}

	bci := Parse(abspath)
	if bci.NumServices() < 1 {
		log.Fatalf("There is no service in %s", bci.PkgName)
	}

	Output(os.Stdout, bci)
}
