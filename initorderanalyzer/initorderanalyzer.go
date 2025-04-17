package initorderanalyzer

import (
	"fmt"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "initorderanalyzer",
	Doc:  "パッケージ初期化順の取得",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, initializer := range pass.TypesInfo.InitOrder {
		fmt.Println("InitOrder:")
		for _, obj := range initializer.Lhs {
			fmt.Printf("初期化対象: %s (%v)\n", obj.Name(), obj.Type())
		}
		fmt.Printf("初期化式: %v\n", initializer.Rhs)
	}
	return nil, nil
}
