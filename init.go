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

	// 忽略标准库
	ignoreStdlib = flag.Bool("nostdlib", false, "ignore packages in the Go standard library")
	// 忽略 vendor 目录下的内容
	ignoreVendor = flag.Bool("novendor", false, "ignore packages in the vendor directory")
	// 当模块引入其他包时，发生错误，就停止
	stopOnError = flag.Bool("stoponerror", false, "stop on package import errors")
	// TODO: 显示标准库中的依赖
	withGoroot = flag.Bool("withgoroot", false, "show dependencies of packages in the Go standard library")
	// 想要忽略的前缀
	ignorePrefixes = flag.String("ignoreprefixes", "", "a comma-separated list of prefixes to ignore")
	// 想要忽略的包
	ignorePackages = flag.String("ignorepackages", "", "a comma-separated list of packages to ignore")
	// 只追踪这些前缀的库
	onlyPrefix = flag.String("onlyprefixes", "", "a comma-separated list of prefixes to include")
	// TODO: 只追踪这些包含这些 tag 的内容
	tagList = flag.String("tags", "", "a comma-separated list of build tags to consider satisfied during the build")
	// 水平显示图像
	horizontal = flag.Bool("horizontal", false, "lay out the dependency graph horizontally instead of vertically")
	// 包括测试包
	withTests = flag.Bool("withtests", false, "include test packages")
	// 依赖图的最大层级
	maxLevel = flag.Int("maxlevel", 256, "max level of go dependency graph")

	// TODO: 什么意思呢
	buildTags    []string
	buildContext = build.Default
)

func init() {
	flag.BoolVar(ignoreStdlib, "s", false, "(alias for -nostdlib) ignore packages in the Go standard library")
	flag.StringVar(ignorePrefixes, "p", "", "(alias for -ignoreprefixes) a comma-separated list of prefixes to ignore")
	flag.StringVar(ignorePackages, "i", "", "(alias for -ignorepackages) a comma-separated list of packages to ignore")
	flag.StringVar(onlyPrefix, "o", "", "(alias for -onlyprefixes) a comma-separated list of prefixes to include")
	flag.BoolVar(withTests, "t", false, "(alias for -withtests) include test packages")
	flag.IntVar(maxLevel, "l", 256, "(alias for -maxlevel) maximum level of the go dependency graph")
	flag.BoolVar(withGoroot, "d", false, "(alias for -withgoroot) show dependencies of packages in the Go standard library")
}
