package main

import (
	"io"

	"github.com/segmentio/ksuid"
)

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	Name string
	Type string
}

func generateCode(dest io.Writer, pkg string, structs map[string][]structField) error {
	tmplContext := getInitialCtx(pkg)
	tmplContext.Structs = makeStructs(structs)

	WriteCsvTemplate(dest, tmplContext)
	return nil
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
				Name: fd.Name,
				Type: fd.Type,
			}
			st.Fields = append(st.Fields, f)
		}

		sts = append(sts, st)
	}

	return sts
}

func getInitialCtx(pkg string) Context {
	return Context{
		Pkg:  pkg,
		Uuid: ksuid.New().String(),
	}
}
