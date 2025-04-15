package expranalyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "expranalyzer",
	Doc: "式の型と定数値を表示する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		expr, ok := n.(ast.Expr)
		if !ok {
			return
		}

		if tv, ok := pass.TypesInfo.Types[expr]; ok {
			pos := pass.Fset.Position(expr.Pos())
			fmt.Printf("[%s] 式: %T, 型: %s", pos, expr, tv.Type)
			if tv.Value != nil {
				fmt.Printf(", 値: %s", tv.Value)
			}
			fmt.Println()
		}
	})

	return nil, nil
}