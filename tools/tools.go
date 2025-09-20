//go:build tools

package tools

//go:generate go install github.com/skulidropek/GoLint/cmd/go-lint@13befa895e15e52dbffec7bf827fa69bb5fcf364
//go:generate go install github.com/skulidropek/GoSuggestMembersAnalyzer/cmd/smbgo@latest
//go:generate go install github.com/skulidropek/gotrace/cmd/gotrace-instrument@afda3736f26d21bb0ad41a341e96bee5990f7822

import (
	_ "github.com/skulidropek/GoLint/cmd/go-lint"
	_ "github.com/skulidropek/GoSuggestMembersAnalyzer/smbgo"
	_ "github.com/skulidropek/gotrace/cmd/gotrace-instrument"
)
