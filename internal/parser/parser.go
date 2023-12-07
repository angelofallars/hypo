package parser

import (
	"strings"

	"golang.org/x/net/html"
)

// Parse parses a string into an HTML node.
func Parse(s string) (*html.Node, error) {
	node, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return nil, err
	}

	node = retrieveElement(node)
	return node, nil
}

// retrieveElement extracts the first element we want from the parsed HTML,
// as the initial output node is a root node.
func retrieveElement(node *html.Node) *html.Node {
	if node.Type == html.DocumentNode && node.Data == "" {
		//    <?> =><html>   =><head>   =><body>    =><[elem]>
		return node.FirstChild.FirstChild.NextSibling.FirstChild
	}
	return node
}
