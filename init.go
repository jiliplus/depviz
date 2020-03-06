package main

import (
	"flag"
	"fmt"
	"go/build"
	"log"
	"os"
	"strings"
)

var (
	// 用于 dfs 保存收集到的模块
	// 避免重复添加
	pkgs = map[string]*pkg{}

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

	interestedStr = flag.String("interested", "", "a comma-separated list of interested directories")
	isInterested  map[string]bool

	deleteStr = flag.String("delete", "", "a comma-separated list of name of want delete directories")
	isDeleted map[string]bool

	showShortName bool
)

func init() {
	flag.BoolVar(help, "h", false, "alias for -help")
	flag.BoolVar(showStandardLib, "s", false, "alias for -showStandard")
	flag.BoolVar(showOtherLib, "o", false, "alias for -showOther")
	flag.BoolVar(&showShortName, "short", false, "short import path, just show directory name")

	flag.StringVar(interestedStr, "i", "", "alias for -interested")
	flag.StringVar(deleteStr, "d", "", "alias for -delete")
}

func prepare() {
	flag.Parse()

	isInterested = makeDict(interestedStr)
	isInterested[getPath()] = true

	isDeleted = makeDict(deleteStr)

	if *help || len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}
}

func usage() {
	// TODO: 添加版本功能
	fmt.Println(`Usage:depviz [options] [package import path | relative path]
	
Options:`)
	flag.PrintDefaults()
}

func makeDict(sp *string) map[string]bool {
	res := make(map[string]bool, 64)
	ss := strings.Split(*sp, ",")
	for _, s := range ss {
		res[s] = true
	}
	return res
}

func getPath() string {
	args := flag.Args()
	if len(args) != 1 {
		log.Fatal("depviz need exactly one path to work")
	}
	// check path
	path := args[0]
	//
	wd := ""
	var err error
	if strings.HasPrefix(path, ".") {
		wd, err = os.Getwd()
		if err != nil {
			log.Fatalf("failed to get cwd: %s", err)
		}
		// log.Printf("current work directory: %s\n", wd)
	}
	// 按照输入参数导入 p
	p, err := build.Import(path, wd, 0)
	if err != nil {
		log.Fatal("can not import with your path, err: ", err)
	}
	// 转换成 import path
	return p.ImportPath
}
