package main

import (
	"go/parser"
    "go/token"
    "go/types"
    "go/ast"
    "log"
)

func main() {
	src := `
		package main
		import "fmt"

		func main() {
			var x int
			fmt.Println(x + "hello") // 型エラー
		}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", src, 0)
	if err != nil {
		log.Fatal(err)
	}

	conf := types.Config{Importer: nil, Error: func(err error) {
		log.Println("型エラー:", err)
	}}

	// 型情報を格納する用の構造体
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	_, err = conf.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Println("チェック中にエラーが発生しました:", err)
	}
}