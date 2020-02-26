package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	fmt.Println(len(os.Args), os.Args)

	fmt.Printf("show standard library: %t\n", *showStandardLib)
	fmt.Printf("show other library   : %t\n", *showOtherLib)
	fmt.Printf("help                 : %t\n", *help)
	fmt.Printf("short path           : %t\n", shortPath)
	fmt.Println()
	fmt.Printf("interested dir = %v\n", *interestedDirs)
	fmt.Printf("delete dir = %v\n", *deleteDirs)

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

// 	// g := gographviz.NewEscape()
// 	// g.SetName("G")
// 	// g.SetDir(false)
// 	// g.SetStrict(true)
// 	// g.AddNode("G", "node", map[string]string{
// 	// 	"color": "black",
// 	// })
// 	// g.AddSubGraph("G", "cluster0", map[string]string{"label": "root"})
// 	// g.AddSubGraph("cluster0", "cluster_1", map[string]string{"label": "child 1"})
// 	// g.AddSubGraph("cluster0", "cluster_2", map[string]string{"label": "child 2"})
// 	// nodeStyle := map[string]string{"style": "\"rounded,filled\"", "color": "\"#00eeee\""}
// 	// nodeStyle2 := map[string]string{"style": "\"rounded,filled\"", "color": "\"#66cd00\""}
// 	// nodeStyle3 := map[string]string{"style": "\"rounded,filled\"", "color": "\"#cdad00\""}
// 	// g.AddNode("cluster_1", "1", nodeStyle)
// 	// g.AddNode("cluster_1", "2", nodeStyle2)
// 	// g.AddNode("cluster_2", "3", nodeStyle3)
// 	// g.AddNode("cluster_2", "4", nodeStyle)
// 	// g.AddEdge("2", "3", false, nil)
// 	// g.AddNode("G", "Code deployment", map[string]string{"style": "dotted"})
// 	// g.AddPortEdge("cluster_2", "", "cluster_1", "", false, nil)
// 	// s := g.String()
// 	// fmt.Println(s)

// 	graph, err := gographviz.Read([]byte(`digraph G {hello->world}`))
// 	if err != nil {
// 		panic(err)
// 	}

// 	// gographviz.new

// 	graph.AddSubGraph("G", "subG", nil)
// 	graph.AddAttr("subG", "color", "red")
// 	//
// 	graph.AddNode("G", "hello", nil)
// 	graph.AddNode("G", "world", nil)
// 	graph.AddEdge("hello", "world", true, nil)
// 	//
// 	thirdPkgNodeAttrs := map[string]string{
// 		"shape":     "box",
// 		"style":     "\"rounded,filled,bold\"",
// 		"fillcolor": "\"#40e0d0\"",
// 		"fontsize":  "24",
// 	}
// 	//
// 	edgeAttrs := map[string]string{
// 		"arrowhead": "open",
// 	}
// 	//
// 	graph.AddAttr("G", "splines", "spline")
// 	graph.AddAttr("G", "center", "true")
// 	// graph.AddAttr("G", "nodesep", "0.4")
// 	// graph.AddAttr("G", "ranksep", "0.8")
// 	//
// 	graph.AddNode("G", "a", nil)
// 	graph.AddNode("G", "subA", nil)
// 	graph.AddNode("G", "subB", nil)
// 	// graph.AddEdge("subA", "subB", true, edgeAttrs)
// 	graph.AddNode("G", "b", thirdPkgNodeAttrs)
// 	graph.AddNode("G", "第三方模块", thirdPkgNodeAttrs)
// 	graph.AddEdge("a", "b", true, edgeAttrs)
// 	graph.AddEdge("world", "a", true, nil)
// 	graph.AddEdge("hello", "a", true, nil)
// 	graph.AddEdge("world", "b", true, nil)
// 	graph.AddEdge("hello", "b", true, nil)
// 	graph.AddEdge("hello", "第三方模块", true, edgeAttrs)
// 	output := graph.String()
// 	fmt.Println(output)

// 	// subpkg.HelloWorld()

// 	// cwd, err := os.Getwd()
// 	// if err != nil {
// 	// 	log.Fatalf("failed to get cwd: %s", err)
// 	// }
// 	// fmt.Printf("current work directory: %s\n", cwd)

// 	// pkg, err := build.Import("github.com/bxcodec/go-clean-arch", "", 0)
// 	// pkg, err := build.Import("github.com/jujili/depviz", "", 0)
// 	// pkg, err := build.Import("./subpkg", cwd, 0)
// 	// pkg, err := build.Import("./subpkg", "", 0)
// 	// pkg, err := build.ImportDir(cwd, 0)
// 	// pkg, err := build.ImportDir(cwd+"/subpkg", 0)
// 	// pkg, err := build.ImportDir(cwd+"/.", 0)

// 	// pkg, err := build.Import(".", cwd, 0)
// 	// if err != nil {
// 	// 	log.Fatal("build.Import err:", err)
// 	// }
// 	// fmt.Printf("imports, %v\n", pkg.Imports)
// 	// fmt.Printf("import path, %v\n", pkg.ImportPath)

// // }
