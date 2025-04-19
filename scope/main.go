package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"analysis/scopetest1"
	// "analysis/scopetest"
)


func main() {
	singlechecker.Main(scopetest.Analyzer)
}