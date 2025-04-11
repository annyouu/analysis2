package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

// evalは式文字列を受け取り、float64で結果を返す
func eval(src string) (float64, error) {
	// パース
	expr, err := parser.ParseExpr(src)
	if err != nil {
		return 0, err
	}

	// 評価
	return evalExpr(expr), nil
}

func evalExpr(e ast.Expr) float64 {
	switch v := e.(type) {
	case *ast.BasicLit:
		// 数値リテラル
		f, _ := strconv.ParseFloat(v.Value, 64)
		return f
	
	case *ast.ParenExpr:
		// かっこ付き式
		return evalExpr(v.X)
	
	case *ast.BinaryExpr:
		// 二項演算子
		left := evalExpr(v.X)
		right := evalExpr(v.Y)

		switch v.Op {
		case token.ADD: // +
			return left + right
		case token.SUB: // -
			return left - right
		case token.MUL: // *
			return left * right
		case token.QUO: // /
			return left / right
		}
	}
	return 0
}


func main() {
	tests := []string{
		"1+2*3",
		"(1+2)*3",
		"10/ (2 + 3)",
		"4- 2 - 1",
	}

	for _, src := range tests {
		v, err := eval(src)
		if err != nil {
			fmt.Printf("パースエラー: %q: %v\n", src, err)
			continue
		}
		fmt.Printf("%s = %v\n", src, v)
	}
}