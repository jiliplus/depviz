package main

import "github.com/awalterschulze/gographviz"

//
var showStdLib = true

//
// g should have a graph named "G"
func dfs(path, parent string, g *gographviz.Graph) *pkg {
	p, ok := pkgs[path]
	if ok {
		return p
	}
	// 记录新的 pkg
	p = newPkg(path)
	pkgs[path] = p
	// 在 graph 中添加点
	if !p.isON() {
		pkgs[path] = nil
		return nil
	}

	g.AddNode("G", p.name(), nodeAttr[p.kind()])

	// if parent != "" {
	// 	g.AddEdge(parent, p.name(), true, edgeAttrs)
	// }

	// 如果 p 是标准库的话，就此打住
	if p.isStandLib() {
		return p
	}

	for _, path := range p.Imports {
		// p.deps = append(p.deps, dfs(path, p.name(), g))
		child := dfs(path, p.name(), g)
		if !child.isON() {
			continue
		}
		g.AddEdge(p.name(), child.name(), true, edgeAttrs)
	}

	return p
}
