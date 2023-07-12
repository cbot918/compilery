package main

import (
	"log"
)

const (
	file = "poc/1.jsx"
)

func main() {
	c, err := NewHtmlJsxCompiler(file)
	if err != nil {
		log.Fatal(err)
	}
	lf("\ncontent: %s\n", string(c.Content))

	tokens, _ := c.Tokenize(c.Content)
	lf("\ntokens: %+v\n", tokens)

	node, _ := c.Noder(tokens)
	lf("\nnode: %+v\n", node)

	result, _ := c.Translate(node)
	lf("\n%s\n\n", result)
	// c.Parse(node)
	// lf("node: %+v\n", node)
}
