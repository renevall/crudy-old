package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "./model", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, ListStruct)
		}
	}
}

//ListStruct lists the name of the structs inside the node package
func ListStruct(node ast.Node) bool {
	file, ok := node.(*ast.File)
	if !ok {
		return false
	}
	for _, d := range file.Decls {
		st, ok := d.(*ast.GenDecl)
		if !ok {
			continue
		}
		if st.Tok.String() == "type" {
			for _, spec := range st.Specs {
				st, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				fmt.Println("Struct found:", st.Name.Name)
				GenCRUD("store", st.Name.Name)
			}
		}
	}
	return true
}
