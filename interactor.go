package main

//
var showStdLib = true

//
func dfs(path string, pkgs map[string]*pkg) *pkg {
	p, ok := pkgs[path]
	if ok {
		return p
	}
	p = newPkg(path)
	pkgs[path] = p
	// 如果 p 是标准库的话，就此打住
	if p.Goroot {
		return p
	}
	for _, path := range p.Imports {
		p.deps = append(p.deps, dfs(path, pkgs))
	}
	return p
}
