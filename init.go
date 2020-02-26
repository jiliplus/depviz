package main

import (
	"flag"
	"go/build"
)

var (
	// 收集模块
	pkgs map[string]*build.Package
	// 错误的模块
	erroredPkgs map[string]bool
	// TODO: ids 是什么意思
	ids map[string]string

	// 忽略的模块
	ignored = map[string]bool{
		"C": true,
	}

	// 忽略的前缀
	ignoredPrefixes []string
	// 只考虑这些前缀
	onlyPrefixes []string

	showStandardLib = flag.Bool("showStandard", false, "show packages in the GOROOT directory")
	help            = flag.Bool("help", false, "show help")
	showOtherLib    = flag.Bool("showOther", false, "show other packages")
	interestedDirs  = flag.String("interested", "", "a comma-separated list of interested directories")
	deleteDirs      = flag.String("delete", "", "a comma-separated list of name of want delete directories")

	// TODO: 支持这个功能，只追踪这些包含这些 tag 的内容
	// tagList = flag.String("tags", "", "a comma-separated list of build tags to consider satisfied during the build")

	shortPath bool
)

func init() {
	flag.BoolVar(showStandardLib, "s", false, "alias for -showStandard")
	flag.BoolVar(help, "h", false, "alias for -help")
	flag.BoolVar(showOtherLib, "o", false, "alias for -showOther")
	flag.BoolVar(&shortPath, "short", false, "short import path, just show directory name")

	flag.StringVar(interestedDirs, "i", "", "alias for -interested")
	flag.StringVar(deleteDirs, "d", "", "alias for -delete")
}
