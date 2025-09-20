package main

import (
	analysis "golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/printf"

	"github.com/skulidropek/GoSuggestMembersAnalyzer/smbgo"
	"honnef.co/go/tools/staticcheck"
)

func main() {
	// bootstrap the builtin analyzers before augmenting with third-party checks
	var analyzers []*analysis.Analyzer
	analyzers = append(analyzers,
		assign.Analyzer,
		atomic.Analyzer,
		bools.Analyzer,
		copylock.Analyzer,
		printf.Analyzer,
	)

	// extend with the typo suggestion analyzer bundled with the devkit
	analyzers = append(analyzers, smbgo.Analyzer)

	// finally add the staticcheck suite for broad bug-finding coverage
	for _, a := range staticcheck.Analyzers {
		analyzers = append(analyzers, a.Analyzer)
	}

	multichecker.Main(analyzers...)
}
