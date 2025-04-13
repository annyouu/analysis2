package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
)

func main() {
	fset := token.NewFileSet()

	srcExpr := `int(1) + 2`

	expr, err := parser.ParseExprFrom(fset, "expr.go", srcExpr, 0)
	if err != nil {
		fmt.Printf("ParseExpr error: %v\n", err)
		return
	}

	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
	}

	if err := types.CheckExpr(fset, nil, token.NoPos, expr, info); err != nil {
		fmt.Printf("Type check error: %v\n", err)
		return
	}

	// 結果を表示
	fmt.Println("式ごとの型と定数値:")
	for e, tv := range info.Types {
		// ASTノードを元のソースに近い形で出力するためのWriteExprの使用
		var buf bytes.Buffer
		types.WriteExpr(&buf, e)
		fmt.Printf(" %s : 型=%s", buf.String(), tv.Type)
		if tv.Value != nil {
			fmt.Printf(", 値=%s", tv.Value)
		}
		fmt.Println()
	}
}