package evaluator

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/angelofallars/hypo/internal/object"
	"golang.org/x/net/html"
)

type command func(node *html.Node, env *object.Env) error

type number = float64

type binOp uint

const (
	add binOp = iota
	subtract
	multiply
	divide
)

var ErrCmdNotFound = errors.New("could not find command definition for this tag")

// Exec executes an HTML, the programming language, program
// from an [html.Node], traversing through it and its siblings.
func Exec(node *html.Node, env *object.Env) error {
	for n := node; n != nil; n = n.NextSibling {
		err := eval(n, env)
		if err != nil {
			fmt.Printf("error on <%v>: %v\n", n.Data, err)
		}
	}

	return nil
}

// eval evaluates a single [html.Node].
func eval(node *html.Node, env *object.Env) error {
	// We are only interested in element nodes, which
	// are the only statements in HTML, the programming language
	if node.Type != html.ElementNode {
		return nil
	}

	cmd, ok := commands[node.Data]
	if !ok {
		return ErrCmdNotFound
	}

	err := cmd(node, env)
	if err != nil {
		return err
	}

	return nil
}

// commands maps an HTML tag name to a command function.
var commands = map[string]command{
	// ===============================
	// Literals
	// ===============================
	"s":    evalPushString,
	"data": evalPushNumber,
	// TODO: "ol"
	// TODO: "table"

	// ===============================
	// Math Commands
	// ===============================
	"dd":  evalBinOp(add),
	"sub": evalBinOp(subtract),
	"ul":  evalBinOp(multiply),
	"div": evalBinOp(divide),

	// ===============================
	// Stack Manipulation Commands
	// ===============================
	"dt":  evalDuplicate,
	"del": evalDelete,

	// ===============================
	// Comparison Commands
	// ===============================
	// TODO: "big"
	// TODO: "small"
	// TODO: "em"

	// ===============================
	// Logical Operators
	// ===============================
	// TODO: "b"
	// TODO: "bdi"
	// TODO: "bdo"

	// ===============================
	// Control Flow
	// ===============================
	// TODO: "i"
	// TODO: "rt"
	// TODO: "a"

	// ===============================
	// Variables
	// ===============================
	// TODO: "var"
	// TODO: "cite"

	// ===============================
	// I/O
	// ===============================
	// TODO: "input"
	"output": evalPrintOutput,
	// TODO: "wbr"

	// ===============================
	// Properties
	// ===============================
	// TODO: "rp"
	// TODO: "samp"

	// ===============================
	// Arrays/Dynamic Properties
	// ===============================
	// TODO: "address"
	// TODO: "ins"

	// ===============================
	// Functions
	// ===============================
	// TODO: "dfn"

	// ===============================
	// Programs
	// ===============================
	// TODO: "main"
	// TODO: "body"
}

// evalPushString pushes a string into the stack.
func evalPushString(node *html.Node, env *object.Env) error {
	obj := &object.String{Value: node.FirstChild.Data}
	env.Stack.Push(obj)
	return nil
}

// evalPushNumber pushes a number into the stack.
func evalPushNumber(node *html.Node, env *object.Env) error {
	attrs := attrMap(node)

	value, ok := getAttr(attrs, "value")
	if !ok {
		return fmt.Errorf("attribute 'value' not found")
	}

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("value is not a valid number")
	}

	obj := &object.Number{Value: number}
	env.Stack.Push(obj)

	return nil
}

// evalBinOp performs a binary operation on the top two values of the stack.
func evalBinOp(op binOp) command {
	return func(node *html.Node, env *object.Env) error {
		top, err := env.Stack.Pop()
		if err != nil {
			return err
		}
		next, err := env.Stack.Pop()
		if err != nil {
			return err
		}

		// TODO: support addition for other types

		topNum, ok := top.(*object.Number)
		if !ok {
			return errors.New("first value is not a number")
		}

		nextNum, ok := next.(*object.Number)
		if !ok {
			return errors.New("second value is not a number")
		}

		// TODO: use a method for the object for arithmetic

		var result number
		switch op {
		case add:
			result = topNum.Value + nextNum.Value
		case subtract:
			result = topNum.Value - nextNum.Value
		case multiply:
			result = topNum.Value * nextNum.Value
		case divide:
			result = topNum.Value / nextNum.Value
		}

		obj := &object.Number{Value: result}
		env.Stack.Push(obj)

		return nil
	}
}

// evalDuplicate duplicates the top value on the stack.
func evalDuplicate(node *html.Node, env *object.Env) error {
	v, err := env.Stack.Top()
	if err != nil {
		return err
	}

	env.Stack.Push(v)
	return nil
}

// evalDelete deletes the top value on the stack.
func evalDelete(node *html.Node, env *object.Env) error {
	_, err := env.Stack.Pop()
	return err
}

// evalPrintOutput prints the top value without consuming it to stdout.
func evalPrintOutput(node *html.Node, env *object.Env) error {
	last, err := env.Stack.Top()
	if err != nil {
		return err
	}

	fmt.Println(last.Display())
	return nil
}

// attrMap creates a map from the Attr slice of an [html.Node].
func attrMap(node *html.Node) map[string]string {
	m := make(map[string]string)
	for _, attr := range node.Attr {
		m[attr.Key] = attr.Val
	}
	return m
}

// getAttr gets a value from the attribute map, if it exists.
func getAttr(attrs map[string]string, key string) (string, bool) {
	value, ok := attrs[key]
	if !ok {
		return "", false
	}
	return value, true
}
