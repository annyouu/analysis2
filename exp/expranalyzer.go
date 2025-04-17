package expranalyzer

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "e",
	Doc: "式から型情報を取得する",
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
		typ := pass.TypesInfo.TypeOf(expr)
		if typ == nil {
			return
		}

		var buf bytes.Buffer
		types.WriteExpr(&buf, expr)
		fmt.Printf("式: %s, 型: %s\n", buf.String(), typ.String())
	})
	return nil, nil
}