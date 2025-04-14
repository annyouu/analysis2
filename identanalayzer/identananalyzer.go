package identanalyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect" 
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "identanalyzer",
	Doc:      "識別子の定義と使用を表示する",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectorResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{(*ast.Ident)(nil)}

	// Preorder走査
	inspectorResult.Preorder(nodeFilter, func(n ast.Node) {
		ident, ok := n.(*ast.Ident)
		if !ok {
			return 
		}

		// 位置情報
        pos := pass.Fset.Position(ident.Pos())
        fmt.Printf("識別子: %s (%s)\n", ident.Name, pos)

        // 定義
        if def := pass.TypesInfo.Defs[ident]; def != nil {
            fmt.Printf("  定義: %s (%s)\n", def.Name(), def.Type())
        }

        // 使用
        if use := pass.TypesInfo.Uses[ident]; use != nil {
            fmt.Printf("  使用: %s (%s)\n", use.Name(), use.Type())
        }
	})

	return nil, nil
}