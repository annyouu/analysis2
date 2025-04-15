package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	src := `
		package main

		func main() {
			var a int
			var b int
			var c string

			_, _, _ = a, b, c
		}
	`

	// ソースコードをパースしてASTを生成する
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "sample.go", src, 0)
	if err != nil {
		panic(err)
	}

	// 型情報を構築
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs: make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("main", fset, []*ast.File{file}, info)
	if err != nil {
		panic(err)
	}

	// 変数a,b,cの型を取得して比較する
	var objA, objB, objC types.Object
	for ident, obj := range info.Defs {
		if ident.Name == "a" {
			objA = obj
		}
		if ident.Name == "b" {
			objB = obj
		}
		if ident.Name == "c" {
			objC = obj
		}
	}

	// 型比較
	if types.Identical(objA.Type(), objB.Type()) {
		fmt.Println("aとbは同じ型です")
	} else {
		fmt.Println("aとbは異なる型")
	}

	if types.Identical(objA.Type(), objC.Type()) {
		fmt.Println("aとcは同じ型")
	} else {
		fmt.Println("aとcは異なる型")
	}
}