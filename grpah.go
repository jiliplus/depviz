package main

var nodeAttr = map[kind]map[string]string{
	interested: interestedNode,
	standard:   standardNode,
	other:      otherNode,
}

var interestedNode = map[string]string{
	"shape":     "box",
	"style":     "\"rounded,filled,bold\"",
	"fillcolor": "\"#40e0d0\"",
	"fontsize":  "24",
}

var standardNode = map[string]string{
	"shape":     "ellipse",
	"style":     "filled",
	"fillcolor": "\"#66cd00\"",
	"fontsize":  "18",
}

var otherNode = map[string]string{
	"shape":     "box",
	"style":     "\"rounded,filled\"",
	"fillcolor": "\"#c0c0c0\"",
	"fontsize":  "21",
}

var edgeAttrs = map[string]string{
	// "arrowhead": "open",
}
