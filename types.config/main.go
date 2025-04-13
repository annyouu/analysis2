package main

import (
	"fmt"
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"go/importer"
	"os"
	"go/printer"
)

func main() {
	// 対象のGoファイル名
	src := `
		package main

		import "fmt"

		func main() {
			var x int = 10
			var y = x + 5
			fmt.Println(y)
		}
	`

	// ファイルセットを作成
	fs := token.NewFileSet()

	// GoファイルをASTにパース
	f, err := parser.ParseFile(fs, "sample.go" ,src, parser.AllErrors)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse error: %v\n", err)
		os.Exit(1)
	}

	// 型チェック結果を保持するためのInfo構造体を初期化する
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
		Implicits: make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes: make(map[ast.Node]*types.Scope),
		InitOrder: nil, // 初期化順は必要な場合のみ
	}

	// Configを設定する (型チェックのオプション)
	cfg := &types.Config{
		Importer: importer.Default(),
		Error: func(err error) {
			fmt.Println("Type error:", err)
		},
	}

	// 型チェックを実行
	pkg, err := cfg.Check("main", fs, []*ast.File{f}, info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Type checking failed: %v\n", err)
	}

	fmt.Println("パッケージ名:", pkg.Name())

	// 変数定義とその型を表示
	fmt.Println("定義されている識別子とその型:")
	for ident, obj := range info.Defs {
		if ident != nil && obj != nil {
			fmt.Printf("- %s: %s\n", ident.Name, obj.Type())
		}
	}

	// 式とその型情報を表示
	fmt.Println("\n式とその型情報:")
	for expr, tv := range info.Types {
		var buf bytes.Buffer
		err := printer.Fprint(&buf, fs, expr)
		if err != nil {
			continue
		}
		pos := fs.Position(expr.Pos())
		fmt.Printf("- %s: %s (型: %s)\n", pos, buf.String(), tv.Type)
	}
}