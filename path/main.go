package main

import (
	"fmt"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"
)

func main() {
	const src = `
		package main

		var v = 100
		func main() {
			fmt.Println(v+1)
		}	
	`

	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "my.go", src, 0)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return 
	}

	fset.Iterate(func(f *token.File) bool {
		if f.Name() != "my.go" {
			return true
		}

		pos := fset.File(file.Pos()).Pos(10)

		// ASTノードを内側 → 外側でリストアップ
		path, exact := astutil.PathEnclosingInterval(file, pos, pos)
		if exact {
			for _, n := range path {
				fmt.Printf("%T\n", n)
			}
		}
		return true
	})
}