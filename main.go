package main

import (
	"fmt"

	"github.com/awalterschulze/gographviz"
)

func main() {
	prepare()

	g, _ := gographviz.Read([]byte(`digraph G {}`))
	g.AddAttr("G", "splines", "ortho")
	path := getPath()

	dfs(path, "", g)

	fmt.Println(g)
}
