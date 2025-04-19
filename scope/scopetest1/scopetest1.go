package scopetest

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"go/types"
)

var Analyzer = &analysis.Analyzer{
	Name: "scopetest",
	Doc: "スコープからオブジェクトを取得する",
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	inspect.Preorder(nil, func(n ast.Node) {
		s := pass.TypesInfo.Scopes[n]
		if s == nil {
			return
		}

		// スコープ内でgopherという名前のオブジェクトを探す
		if obj := s.Lookup("gopher"); obj != nil {
			fmt.Println("オブジェクト:", obj)
		}

		// 親スコープを辿ってgopherを探す
		if obj := lookupParentScope(s, "gopher"); obj != nil {
			fmt.Print("親スコープから辿って:", obj)
		}
	})
	return nil, nil
}

func lookupParentScope(s *types.Scope, name string) types.Object {
	for parent := s.Parent(); parent != nil; parent = parent.Parent() {
		if obj := parent.Lookup(name); obj != nil {
			return obj
		}
	}
	return nil
}