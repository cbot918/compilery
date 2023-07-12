package main

import (
	"fmt"
	"os"
)

// type Compiler interface {
// 	Read(file string) (string, error)
// 	Tokenize(content string) (*Token, error)
// 	Parse()
// 	Executer()
// }

type HtmlJsxCompiler struct {
	Content            []byte
	Index              int
	Length             int
	timeToGetAttribute bool
	timeToGetInnerText bool
	Tokens             []Token
}

func NewHtmlJsxCompiler(path string) (*HtmlJsxCompiler, error) {

	c := new(HtmlJsxCompiler)
	content, err := c.Read(path)
	if err != nil {
		return nil, err
	}

	return &HtmlJsxCompiler{
		Content:            content,
		Index:              0,
		Length:             len(content),
		timeToGetAttribute: false,
		timeToGetInnerText: false,
	}, nil
}

func (h *HtmlJsxCompiler) Read(file string) ([]byte, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		return []byte{}, err
	}
	return content, nil
}

// temp return One node
func (h *HtmlJsxCompiler) Noder(tokens []Token) (Node, error) {
	node := Node{}
	for _, ele := range tokens {
		if ele.Symbol == "TAG" {
			node.Tag = ele.Value
		}
		if ele.Symbol == "ATTR" {
			node.Attr = ele.Value
		}
		if ele.Symbol == "INNERTEXT" {
			node.Child = ele.Value
		}
	}
	return node, nil
}

func (h *HtmlJsxCompiler) Parse(node Node) {

}
func (h *HtmlJsxCompiler) Translate(node Node) (string, error) {
	return fmt.Sprintf("('%s','%s','%s')", node.Tag, node.Attr, node.Child), nil
}
