package main

import (
	"io"
	"strings"

	"github.com/knadh/stuffbin"
)

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	Name      string
	NameLower string
	Type      string
}

func generateCode(dest io.Writer, fs stuffbin.FileSystem, pkg string, structs map[string][]structField) error {
	tmplContext := make(map[string]interface{})
	addInitialContext(tmplContext, pkg)

	tmplContext["Structs"] = makeStructs(structs)

	return saveResource("struct", []string{"/templates/gen.tmpl"}, dest, tmplContext, fs)
}

func makeStructs(structs map[string][]structField) []Struct {
	sts := make([]Struct, 0, len(structs))
	for name, fields := range structs {
		st := Struct{
			StructName: name,
			Fields:     make([]Field, 0, len(fields)),
		}

		for _, fd := range fields {
			f := Field{
				Name:      fd.Name,
				NameLower: strings.ToLower(fd.Name),
				Type:      fd.Type,
			}
			st.Fields = append(st.Fields, f)
		}

		sts = append(sts, st)
	}

	return sts
}

func addInitialContext(tmplContext map[string]interface{}, pkg string) {
	tmplContext["Pkg"] = pkg
	tmplContext["BuildDate"] = buildDate
	tmplContext["BuildVersion"] = buildVersion
}
