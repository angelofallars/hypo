package ast

import (
	"fmt"
	"strings"

	"github.com/angelofallars/hypo/pkg/sliceutil"
)

type Node interface {
	// Dummy function
	astNode()
	String() string
}

type Program struct {
	Statements []Node
}

func (p *Program) astNode() {}
func (p *Program) String() string {
	childStrings := sliceutil.Map(p.Statements, func(stmt Node) string { return stmt.String() })
	return fmt.Sprintf("%v", strings.Join(childStrings, "\n"))
}

type NumberStatement struct {
	Value float64
}

func (ns *NumberStatement) astNode() {}
func (ns *NumberStatement) String() string {
	return fmt.Sprintf(`<data value="%f"></data>`, ns.Value)
}

type StringStatement struct {
	Value string
}

func (ss *StringStatement) astNode() {}
func (ss *StringStatement) String() string {
	return fmt.Sprintf(`<s>%s</s>`, ss.Value)
}

type BoolStatement struct {
	Value bool
}

func (bs *BoolStatement) astNode() {}
func (bs *BoolStatement) String() string {
	return fmt.Sprintf(`<cite>%v</cite>`, bs.Value)
}

type ArrayStatement struct {
	Statements []Node
}

func (as *ArrayStatement) astNode() {}
func (as *ArrayStatement) String() string {
	childStrings := sliceutil.Map(as.Statements, func(stmt Node) string { return stmt.String() })
	return fmt.Sprintf(`<ol>%v</ol>`, strings.Join(childStrings, ""))
}

type ArrayElementStatement struct {
	Statements []Node
}

func (aes *ArrayElementStatement) astNode() {}
func (aes *ArrayElementStatement) String() string {
	childStrings := sliceutil.Map(aes.Statements, func(stmt Node) string { return stmt.String() })
	return fmt.Sprintf(`<li>%v</li>`, strings.Join(childStrings, ""))
}

type DuplicateStatement struct{}

func (ds *DuplicateStatement) astNode() {}
func (ds *DuplicateStatement) String() string {
	return "<dup></dup>"
}

type DeleteStatement struct{}

func (ds *DeleteStatement) astNode() {}
func (ds *DeleteStatement) String() string {
	return "<del></del>"
}

type PrintStatement struct{}

func (os *PrintStatement) astNode() {}
func (os *PrintStatement) String() string {
	return "<output></output>"
}

type BinaryOpStatement struct {
	Op BinaryOp
}

func (bos *BinaryOpStatement) astNode() {}
func (bos *BinaryOpStatement) String() string {
	var tag string
	switch bos.Op {
	case BinAdd:
		tag = "dd"
	case BinSubtract:
		tag = "sub"
	case BinMultiply:
		tag = "ul"
	case BinDivide:
		tag = "div"
	default:
		panic(fmt.Sprintf("Binary operation is not recognized: %v", bos.Op))
	}
	return fmt.Sprintf("<%s></%s>", tag, tag)
}

type BinaryOp uint

const (
	BinAdd BinaryOp = iota
	BinSubtract
	BinMultiply
	BinDivide
)

type VariableStatement struct {
	Identifier string
}

func (vs *VariableStatement) astNode() {}
func (vs *VariableStatement) String() string {
	return fmt.Sprintf("<cite>%v</cite>", vs.Identifier)
}
