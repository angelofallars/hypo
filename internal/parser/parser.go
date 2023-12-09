// package parser provides a parser that has two stages:
//
// Stage 1: Parse a [string] into an *[html.Node] tree.
//
// Stage 2: Parse an *[html.Node] tree into an *[ast.Node] tree.
package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/angelofallars/hypo/internal/ast"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Parser transforms the nodes from [html.Parser] into
// validated runtime-specific AST nodes.
type Parser struct {
	curNode  *html.Node
	peekNode *html.Node
}

func New() *Parser {
	return &Parser{
		curNode:  nil,
		peekNode: nil,
	}
}

type parseError struct{ err error }

func newParseError(err error) parseError {
	return parseError{err: err}
}

func (pe parseError) Error() string { return fmt.Sprintf("ParseError: %v", pe.err.Error()) }

// Parse parses a string into Hypo-specific AST nodes.
func Parse(s string) (*ast.Program, error) {
	return New().Parse(s)
}

// Parse parses a string into Hypo-specific AST nodes.
func (p *Parser) Parse(s string) (*ast.Program, error) {
	if err := p.parseString(s); err != nil {
		return nil, newParseError(err)
	}

	program := &ast.Program{
		Statements: []ast.Node{},
	}

	parseErrors := []error{}
	for ; p.curNode != nil; p.nextNode() {
		newNode, err := p.parseStatement()
		if err != nil {
			parseErrors = append(parseErrors, err)
		}
		program.Statements = append(program.Statements, newNode)
	}

	if len(parseErrors) != 0 {
		return nil, newParseError(errors.Join(parseErrors...))
	}

	return program, nil
}

// parseString parses a string into an *[html.Node] tree.
func (p *Parser) parseString(s string) error {
	node, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return err
	}

	//    <?> =><html>   =><head>   =><body>    =><[elem]>
	node = node.FirstChild.FirstChild.NextSibling.FirstChild
	p.peekNode = node
	p.nextNode()
	return nil
}

// nextNode advances the parser's input nodes.
func (p *Parser) nextNode() {
	p.curNode = p.peekNode
	if p.peekNode != nil {
		p.peekNode = p.peekNode.NextSibling
	}

	if p.curNode != nil && p.curNode.Type != html.ElementNode {
		p.nextNode()
	}
}

// parseStatement parses an AST statement from the current [html.Node].
func (p *Parser) parseStatement() (ast.Node, error) {
	var node ast.Node
	var err error

	switch p.curNode.DataAtom {
	case atom.S:
		node, err = p.parseStringStatement()
	case atom.Data:
		node, err = p.parseNumberStatement()
	case atom.Ol:
		node, err = p.parseArrayStatement()
	case atom.Li:
		node, err = p.parseArrayElementStatement()
	case atom.Dd:
		node, err = p.parseBinaryOpStatement(ast.BinAdd)
	case atom.Sub:
		node, err = p.parseBinaryOpStatement(ast.BinSubtract)
	case atom.Ul:
		node, err = p.parseBinaryOpStatement(ast.BinMultiply)
	case atom.Div:
		node, err = p.parseBinaryOpStatement(ast.BinDivide)
	case atom.Dt:
		node, err = p.parseDuplicateStatement()
	case atom.Del:
		node, err = p.parseDeleteStatement()
	case atom.Output:
		node, err = p.parsePrintStatement()
	default:
		err = fmt.Errorf("unknown tag '%v'", p.curNode.Data)
	}

	return node, err
}

func (p *Parser) parseStringStatement() (*ast.StringStatement, error) {
	if p.curNode.FirstChild == nil {
		return nil, errors.New("<s> element has no text child element")
	}

	return &ast.StringStatement{
		Value: p.curNode.FirstChild.Data,
	}, nil
}

func (p *Parser) parseNumberStatement() (*ast.NumberStatement, error) {
	attrs := attrMap(p.curNode)

	value, ok := attrs["value"]
	if !ok {
		return nil, fmt.Errorf("attribute 'value' not found")
	}

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil, errors.New("value is not a valid number")
	}

	return &ast.NumberStatement{
		Value: number,
	}, nil
}

func (p *Parser) parseArrayStatement() (*ast.ArrayStatement, error) {
	array := &ast.ArrayStatement{Statements: []ast.Node{}}

	statements, err := p.parseChildStatements()
	if err != nil {
		return nil, err
	}
	array.Statements = statements

	return array, nil
}

func (p *Parser) parseArrayElementStatement() (*ast.ArrayElementStatement, error) {
	arrayElement := &ast.ArrayElementStatement{Statements: []ast.Node{}}

	statements, err := p.parseChildStatements()
	if err != nil {
		return nil, err
	}
	arrayElement.Statements = statements

	return arrayElement, nil
}

func (p *Parser) parseBinaryOpStatement(binaryOp ast.BinaryOp) (*ast.BinaryOpStatement, error) {
	return &ast.BinaryOpStatement{
		Op: binaryOp,
	}, nil
}

func (p *Parser) parseDuplicateStatement() (*ast.DuplicateStatement, error) {
	return &ast.DuplicateStatement{}, nil
}

func (p *Parser) parseDeleteStatement() (*ast.DeleteStatement, error) {
	return &ast.DeleteStatement{}, nil
}

func (p *Parser) parsePrintStatement() (*ast.PrintStatement, error) {
	return &ast.PrintStatement{}, nil
}

// parseChildStatements parses the child nodes of the current node.
func (p *Parser) parseChildStatements() ([]ast.Node, error) {
	originalNode := p.curNode
	p.peekNode = originalNode.FirstChild
	p.nextNode()

	statements := []ast.Node{}

	parseErrors := []error{}
	for p.curNode != nil {
		newNode, err := p.parseStatement()
		if err != nil {
			parseErrors = append(parseErrors, err)
			continue
		}
		statements = append(statements, newNode)

		p.nextNode()
	}

	p.peekNode = originalNode
	p.nextNode()

	if len(parseErrors) != 0 {
		return nil, errors.Join(parseErrors...)
	}

	return statements, nil
}

func (p *Parser) curNodeIs(atom atom.Atom) bool {
	return p.curNode != nil && p.curNode.DataAtom == atom
}

// attrMap creates a map from the Attr slice of an [html.Node].
func attrMap(node *html.Node) map[string]string {
	m := make(map[string]string)
	for _, attr := range node.Attr {
		m[attr.Key] = attr.Val
	}
	return m
}
