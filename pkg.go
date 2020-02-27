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
		log.Fatal("build.Import err: ", err)
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

func join(dirs []string) string {
	sep := "/"
	if runtime.GOOS == "windows" {
		sep = "\\"
	}
	return strings.Join(dirs, sep)
}

func (p *pkg) isInterestedLib() bool {
	// dirs := strings.Split(p.ImportPath, sep)
	dirs := split(p.ImportPath)
	size := len(dirs)
	for i := 1; i <= size; i++ {
		path := join(dirs[:i])
		if isInterested[path] {
			return true
		}
	}
	return false
}

type kind int

const (
	interested kind = iota
	standard
	other
)

func (p *pkg) kind() kind {
	if p.isStandLib() {
		return standard
	}
	if p.isInterestedLib() {
		return interested
	}
	return other
}

func (p *pkg) shortName() string {
	dirs := split(p.ImportPath)
	size := len(dirs)
	return dirs[size-1]
}

func (p *pkg) isON() bool {
	if p == nil {
		return false
	}
	if p.isInterestedLib() {
		return true
	}
	if p.isStandLib() {
		return *showStandardLib
	}
	return *showOtherLib
}

func (p *pkg) name() string {
	if showShortName {
		return "\"" + p.shortName() + "\""
	}
	dirs := split(p.ImportPath)
	tmp := make([]string, 0, len(dirs))
	for _, d := range dirs {
		if isDeleted[d] {
			continue
		}
		tmp = append(tmp, d)
	}
	return "\"" + join(tmp) + "\""
}
