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

func eval(node *html.Node, env *object.Env) error {
	cmd, ok := commands[node.Data]
	if !ok {
		// TODO: Error for unrecognized tags
		return nil
	}

	err := cmd(node, env)
	if err != nil {
		return err
	}

	return nil
}

// commands maps an HTML tag name to a command function.
var commands = map[string]command{
	// Literals
	"s":    evalPushString,
	"data": evalPushNumber,
	// Math Commands
	"dd":  evalBinOp(add),
	"sub": evalBinOp(subtract),
	"ul":  evalBinOp(multiply),
	"div": evalBinOp(divide),
	// Stack Manipulation commands
	"dt":  evalDuplicate,
	"del": evalDelete,
	// I/O
	"output": evalPrintOutput,
}

// evalPushString pushes a string into the stack.
func evalPushString(node *html.Node, env *object.Env) error {
	env.Stack.Push(node.FirstChild.Data)
	return nil
}

// evalPushNumber pushes a number into the stack.
func evalPushNumber(node *html.Node, env *object.Env) error {
	attrs := attrMap(node)

	value, err := getAttr(attrs, "value")
	if err != nil {
		return err
	}

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return errors.New("value is not a valid number")
	}

	env.Stack.Push(number)

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

		topNum, ok := top.(number)
		if !ok {
			return errors.New("first value is not a number")
		}

		nextNum, ok := next.(number)
		if !ok {
			return errors.New("second value is not a number")
		}

		var result number
		switch op {
		case add:
			result = topNum + nextNum
		case subtract:
			result = topNum - nextNum
		case multiply:
			result = topNum * nextNum
		case divide:
			result = topNum / nextNum
		}

		env.Stack.Push(result)
		return nil
	}
}

// evalPrintOutput prints the top value without consuming it to stdout.
func evalPrintOutput(node *html.Node, env *object.Env) error {
	last, err := env.Stack.Top()
	if err != nil {
		return err
	}

	fmt.Println(last)
	return nil
}

// evalDelete deletes the top value on the stack.
func evalDelete(node *html.Node, env *object.Env) error {
	_, err := env.Stack.Pop()
	return err
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

// attrMap creates a map from the Attr slice of an [html.Node].
func attrMap(node *html.Node) map[string]string {
	m := make(map[string]string)
	for _, attr := range node.Attr {
		m[attr.Key] = attr.Val
	}
	return m
}

// getAttr gets a value from the attribute map, if it exists.
func getAttr(attrs map[string]string, key string) (string, error) {
	value, ok := attrs[key]
	if !ok {
		return "", fmt.Errorf("attribute '%v' not found", key)
	}
	return value, nil
}
