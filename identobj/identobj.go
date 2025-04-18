package identobj

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/passes/inspect"
    "golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "identobj",
	Doc: "識別子からオブジェクトを取得する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{(*ast.Ident)(nil)}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		ident := n.(*ast.Ident)

		obj := pass.TypesInfo.ObjectOf(ident)
		if obj == nil {
			return
		}

		// 位置・名前・型を出力
		pos := pass.Fset.Position(ident.Pos())
		fmt.Printf("位置: %s 識別子: %s オブジェクト: %s (型 %s)\n", pos, ident.Name, obj.Name(), obj.Type())
	})
	return nil, nil
}