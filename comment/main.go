package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	const src = `
		package main

		import "fmt"

		func main() {
			v := 100 // v is int value
			fmt.Println(v + 1)
        }
	`

	// tokenの位置情報を管理する
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, "main.go", src, parser.ParseComments)
	if err != nil {
		fmt.Println("パースエラー:", err)
		return
	}

	// コメントとノードの対応マップを作成する
	cmap := ast.NewCommentMap(fset, file, file.Comments)

	// ASTを走査する
	ast.Inspect(file, func(n ast.Node) bool {
		if assignStmt, ok := n.(*ast.AssignStmt); ok {
			if comments, found := cmap[n]; found {
				fmt.Println("コメント:", comments[0].Text())
			} else {
				fmt.Println("コメントがない代入文です")
			}
			fmt.Printf("代入文: %v\n", assignStmt)
		}
		return true
	})
}