package implicitanalyzer

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "implicitanalyzer",
	Doc: "暗黙に定義された識別子を表示する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectorResult := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspectorResult.Preorder(nil, func(n ast.Node) {
		if obj := pass.TypesInfo.Implicits[n]; obj != nil {
			pos := pass.Fset.Position(n.Pos())
			fmt.Printf("暗黙定義: %T at %s → %s (%s)\n", n, pos, obj.Name(), obj.Type())
		}
	})
	return nil, nil
}