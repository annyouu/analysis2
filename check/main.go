package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"go/importer"
)

func main() {
	const src = `
		package main

		import "fmt"

		func add(a, b int) int {
			return a + b
		}

		func main() {
			x := add(1, 2)
			fmt.Println(x)
		}
	`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example.go", src, parser.AllErrors)
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	// 型チェックの結果を保持する Infoを初期化する
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}

	// 型チェック用のConfigを初期化
	cfg := &types.Config{
		Importer: importer.Default(),
	}

	// 型チェックを実行する
	pkg, err := cfg.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("Type check error: %v\n", err)
	}

	// 型チェック結果を使った処理例
	fmt.Printf("型チェック完了: パッケージ %q\n", pkg.Name())

	// 変数、定数、関数定義の型情報を表示
	fmt.Println("\n=== 定義(Defs) ===")
	for id, obj := range info.Defs {
		if obj != nil {
			fmt.Printf("%s: %s (%s)\n", id.Name, obj.Type(), fset.Position(id.Pos()))
		}
	}

	// 識別子の使用箇所と型を表示
	fmt.Println("\n=== 使用(Uses) ===")
	for id, obj := range info.Uses {
		fmt.Printf("%s -> %s (%s)\n", id.Name, obj.Type(), fset.Position(id.Pos()))
	}

	// 各式の型を表示
	fmt.Println("\n=== 式の型(Types) ===")
	for expr, tv := range info.Types {
		fmt.Printf("%s: %s\n", fset.Position(expr.Pos()), tv.Type)
	}
}

