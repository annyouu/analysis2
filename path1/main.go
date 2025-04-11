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

		import "fmt"

		var v = 100

		func main() {
			fmt.Println(v + 1)
		}
	`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "my.go", src, 0)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return
	}

	astutil.Apply(file, nil, func(cr *astutil.Cursor) bool {
		n :=  cr.Node()
		if n != nil {
			start := fset.Position(n.Pos()).Offset
			end := fset.Position(n.End()).Offset
			fmt.Printf("%T: %d - %d\n", n, start, end)
		}
		return true
	})

	// vのidentノードの範囲がわかったら、その範囲内にposを設定する。
	pos := token.Pos(fset.File(file.Pos()).Base() + 75)

	path, exact := astutil.PathEnclosingInterval(file, pos, pos)
	if !exact {
		fmt.Println("一致するノードがありません")
		return
	}
	fmt.Println("PathEnclosingIntervalの結果")
	for _, n := range path {
		fmt.Printf("%T\n", n)
	}
}