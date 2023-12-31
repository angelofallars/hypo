package evaluator

import (
	"fmt"

	"github.com/angelofallars/hypo/internal/ast"
	errs "github.com/angelofallars/hypo/internal/errors"
	"github.com/angelofallars/hypo/internal/object"
)

// Exec evaluates a single [ast.Node].
func Exec(node ast.Node, env *object.Env) error {
	var err error

	switch node := node.(type) {
	// ===============================
	// Program root
	// ===============================
	case *ast.Program:
		err = evalProgram(node, env)

	// ===============================
	// Literals
	// ===============================
	case *ast.StringStatement:
		err = evalPushString(node, env)
	case *ast.NumberStatement:
		err = evalPushNumber(node, env)
	case *ast.ArrayStatement:
		err = evalPushArray(node, env)
	// case *ast.TableStatement:
	// 	err = evalPushTable(node, env)

	// ===============================
	// Math commands
	// ===============================
	case *ast.BinaryOpStatement:
		err = evalBinOp(node, env)

	// ===============================
	// Stack Manipulation Commands
	// ===============================
	case *ast.DuplicateStatement:
		err = evalDuplicate(node, env)
	case *ast.DeleteStatement:
		err = evalDelete(node, env)

	// ===============================
	// Variables
	// ===============================
	case *ast.SetVariableStatement:
		err = evalSetVariable(node, env)
	case *ast.GetVariableStatement:
		err = evalGetVariable(node, env)

	// ===============================
	// I/O
	// ===============================
	case *ast.PrintStatement:
		err = evalPrint(node, env)
	}

	return err
}

// evalProgram evaluates an [ast.Program].
func evalProgram(program *ast.Program, env *object.Env) error {
	for _, statement := range program.Statements {
		err := Exec(statement, env)
		if err != nil {
			return err
		}
	}
	return nil
}

// evalPushString pushes a string into the stack.
func evalPushString(node *ast.StringStatement, env *object.Env) error {
	object := &object.String{Value: node.Value}
	env.Stack.Push(object)
	return nil
}

// evalPushNumber pushes a number into the stack.
func evalPushNumber(node *ast.NumberStatement, env *object.Env) error {
	object := &object.Number{Value: node.Value}
	env.Stack.Push(object)
	return nil
}

// evalPushArray pushes an array into the stack.
func evalPushArray(node *ast.ArrayStatement, env *object.Env) error {
	obj := &object.Array{}

	elements := []object.Object{}

	initialLength := env.Stack.Len()
	defer func() {
		// Pop any excess objects
		removeCount := env.Stack.Len() - initialLength - 1
		_, _ = env.Stack.PopMany(removeCount)
	}()

	for _, childNode := range node.Elements {
		for _, childChildNode := range childNode.Statements {
			err := Exec(childChildNode, env)
			if err != nil {
				return err
			}
		}

		poppedObject, err := env.Stack.Pop()
		if err != nil {
			return err
		}

		elements = append(elements, poppedObject)

		// Pop any excess objects
		removeCount := env.Stack.Len() - initialLength
		_, _ = env.Stack.PopMany(removeCount)
	}

	obj.Value = elements
	env.Stack.Push(obj)

	return nil
}

// evalBinOp performs a binary operation on the top two values of the stack.
func evalBinOp(node *ast.BinaryOpStatement, env *object.Env) error {
	objects, err := env.Stack.PeekMany(2)
	if err != nil {
		return err
	}

	// Right-hand side is the top-most value in the stack
	right := objects[0]
	left := objects[1]

	switch {
	case left.Type() != right.Type():
		return errs.NewTypeError("cannot perform %v on types '%v' and '%v'",
			node.Op, left.Type(), right.Type())
	case left.Type() == object.NumberType && right.Type() == object.NumberType:
		leftNumber := left.(*object.Number).Value
		rightNumber := right.(*object.Number).Value

		number := 0.0
		switch node.Op {
		case ast.BinAdd:
			number = leftNumber + rightNumber
		case ast.BinSubtract:
			number = leftNumber - rightNumber
		case ast.BinMultiply:
			number = leftNumber * rightNumber
		case ast.BinDivide:
			number = leftNumber / rightNumber
		}

		object := &object.Number{Value: number}

		_, _ = env.Stack.PopMany(2)
		env.Stack.Push(object)
		return nil
	case left.Type() == object.StringType && right.Type() == object.StringType && node.Op == ast.BinAdd:
		leftString := left.(*object.String).Value
		rightString := right.(*object.String).Value

		object := &object.String{Value: leftString + rightString}

		_, _ = env.Stack.PopMany(2)
		env.Stack.Push(object)
		return nil
	default:
		return errs.NewTypeError("cannot perform %v on type '%v'",
			node.Op, left.Type())
	}
}

// evalDuplicate duplicates the top value on the stack.
func evalDuplicate(_ *ast.DuplicateStatement, env *object.Env) error {
	object, err := env.Stack.Peek()
	if err != nil {
		return err
	}

	env.Stack.Push(object)
	return nil
}

// evalDelete deletes the top value on the stack.
func evalDelete(_ *ast.DeleteStatement, env *object.Env) error {
	_, err := env.Stack.Pop()
	return err
}

// evalSetVariable sets a variable in the environment with the top value on the stack.
func evalSetVariable(node *ast.SetVariableStatement, env *object.Env) error {
	object, err := env.Stack.Pop()
	if err != nil {
		return err
	}

	err = env.Vars.Set(node.Identifier, object)
	if err != nil {
		return err
	}

	return nil
}

// evalGetVariable pushes a variable with the given name into the stack.
func evalGetVariable(node *ast.GetVariableStatement, env *object.Env) error {
	object, err := env.Vars.Get(node.Identifier)
	if err != nil {
		return err
	}
	env.Stack.Push(object)
	return nil
}

// evalPrint prints the top value without consuming it to stdout.
func evalPrint(_ *ast.PrintStatement, env *object.Env) error {
	object, err := env.Stack.Peek()
	if err != nil {
		return err
	}
	fmt.Println(object.String())
	return nil
}
