package main

import (
	"github.com/amit-davidson/goclose/passes/goclose"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(goclose.Analyzer)
}
