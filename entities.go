package main

import (
	"go/build"
	"log"
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
