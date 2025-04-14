package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"ana/identanalyzer"
)

func main() {
	singlechecker.Main(identanalyzer.Analyzer)
}


