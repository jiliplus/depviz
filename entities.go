package main

import (
	"go/build"
	"log"
	"runtime"
	"strings"
)

// package 包含了 build.Package
type pkg struct {
	*build.Package
	// dependencies c
	deps []*pkg
}

func newPkg(path string) *pkg {
	p, err := build.Import(path, "", 0)
	if err != nil {
		log.Println("build.Import err: ", err)
	}
	return &pkg{
		Package: p,
	}
}

func (p *pkg) isStandLib() bool {
	return p.Goroot
}

func split(path string) []string {
	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}
	return strings.Split(path, sep)
}

func (p *pkg) isInterestedLib() bool {
	// dirs := strings.Split(p.ImportPath, sep)
	dirs := split(p.ImportPath)
	size := len(dirs)
	for i := 2; i <= size; i++ {
		path := strings.Join(dirs[:i], sep)
		if isInterested[path] {
			return true
		}
	}
	return false
}

type kind int

const (
	interested kind = iota
	stand
	other
)

// TODO: 放入 init 内
var isInterested = make(map[string]bool, 1024)

func (p *pkg) kind() kind {
	if p.isStandLib() {
		return stand
	}
	if p.isStandLib() {
		return interested
	}
	return other
}

func (p *pkg) shortName() string {
	dirs := split(p.ImportPath)
	size := len(dirs)
	return dirs[size-1]
}
