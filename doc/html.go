package doc

import (
	"strings"

	"golang.org/x/net/html"
)

type Node struct {
	ID         string
	Children   []*Node
	Attributes NamedNodeMap
	Type       html.NodeType
	node       *html.Node
}

type HTMLElement struct {
	Name       string
	ID         string
	ClassList  ClassList
	Attributes NamedNodeMap
}

func buildDocument(n *html.Node) *Node {
	var parseNode func(*html.Node) *Node
	parseNode = func(n *html.Node) *Node {
		cl := []*Node{}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			cl = append(cl, parseNode(c))
		}
		return &Node{
			Children: cl,
		}
	}

	return parseNode(n)
}

// ---- Node methods ----

func (n *Node) el() *HTMLElement {
	e := &HTMLElement{
		ID:         n.attr("id"),
		ClassList:  ClassList(strings.Split(n.attr("class"), " ")),
		Attributes: n.Attributes,
	}

	return e
}

func (n *Node) attr(key string) string {
	if val, ok := n.Attributes[key]; ok {
		return val
	}

	return ""
}

func (n *Node) isElement() bool {
	return n.Type == html.ElementNode
}

func (n *Node) firstChild() *Node {
	return n.Children[0]
}

func (n *Node) firstElementChild() *Node {
	for _, c := range n.Children {
		if c.isElement() {
			return c
		}
	}

	return nil
}
