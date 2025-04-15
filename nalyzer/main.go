package main

import (
	"analy/expranalyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(expranalyzer.Analyzer)
}