package main

import (
	"bytes"
	"flag"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
)

var (
	buildDate, buildVersion string
)

func main() {
	pkg := flag.String("pkg", "main", "Package for the generated file")
	fName := flag.String("file", "", "File with structs")
	dest := flag.String("dest", "", "Destination file path to store generated code")
	flag.Parse()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, *fName, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("error while parsing file: %v", err)
	}

	sts, err := parseNode(f)
	if err != nil {
		log.Fatalf("error while parsing ast from the file: %v", err)
	}

	code := bytes.NewBuffer([]byte{})

	if err := generateCode(code, *pkg, sts); err != nil {
		log.Fatalf("error while generating code: %v", err)
	}

	fmted, err := format.Source(code.Bytes())
	if err != nil {
		log.Fatalf("error while formatting code: %v", err)
	}

	genfile, err := os.Create(*dest)
	if err != nil {
		log.Fatalf("error while opening the file for writing code")
	}

	if _, err := genfile.Write(fmted); err != nil {
		log.Fatalf("error while writing code to the file: %v", err)
	}
}
