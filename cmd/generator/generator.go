package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/segmentio/ksuid"
)

type Struct struct {
	StructName string
	Fields     []Field
}

type Field struct {
	Name     string
	Type     string
	Position int
}

func generateCode(dest io.Writer, pkg string, structs map[string][]structField) error {
	tmplContext := getInitialCtx(pkg)
	sts, err := makeStructs(structs)
	if err != nil {
		return err
	}
	tmplContext.Structs = sts

	WriteCsvTemplate(dest, tmplContext)
	return nil
}

func makeStructs(structs map[string][]structField) ([]Struct, error) {
	sts := make([]Struct, 0, len(structs))
	for name, fields := range structs {
		st := Struct{
			StructName: name,
			Fields:     make([]Field, 0, len(fields)),
		}

		pos := 0
		for _, fd := range fields {
			if fd.Tag != "" {
				// Check if there's a tag with position.
				posStr := strings.Trim(strings.ReplaceAll(fd.Tag, "csv:", ""), `"`)
				tagPos, err := strconv.Atoi(posStr)
				if err != nil {
					return nil, fmt.Errorf("invalid struct tag: %v", fd.Tag)
				}

				if tagPos != 0 {
					pos = tagPos
				}
			}

			f := Field{
				Name:     fd.Name,
				Type:     fd.Type,
				Position: pos,
			}
			st.Fields = append(st.Fields, f)
			pos++
		}

		sts = append(sts, st)
	}

	return sts, nil
}

func getInitialCtx(pkg string) Context {
	return Context{
		Pkg:  pkg,
		Uuid: ksuid.New().String(),
	}
}
