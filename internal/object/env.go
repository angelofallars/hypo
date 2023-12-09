package object

import (
	"errors"
	"fmt"
	"slices"
)

// Env is an environment of the runtime which contains runtime values.
type Env struct {
	// Stack is the primary storage of values.
	Stack stack
	Vars  vars
}

// NewEnv returns a new [object.Env] instance.
func NewEnv() *Env {
	return &Env{
		Stack: stack{make([]Object, 0, 256)},
		Vars: vars{
			objects: map[string]Object{
				// Start with standard variables for common values
				"true":  &Bool{Value: true},
				"false": &Bool{Value: false},
				"null":  &Null{},
			},
		},
	}
}

var ErrStackEmpty = errors.New("stack empty")

type stack struct {
	slice []Object
}

// Push a value to the top of the stack.
func (s *stack) Push(v Object) {
	s.slice = append(s.slice, v)
}

// Pop a value from the top of the stack and return it.
func (s *stack) Pop() (Object, error) {
	v, err := s.Top()
	if err != nil {
		return nil, err
	}

	topIndex := s.topIndex()
	s.slice = slices.Delete(s.slice, topIndex, topIndex+1)
	return v, nil
}

// Receive a value from the top of the stack without consuming it.
func (s *stack) Top() (Object, error) {
	if len(s.slice) == 0 {
		return nil, ErrStackEmpty
	}

	return s.slice[s.topIndex()], nil
}

// topIndex returns the last index in the stack.
func (s *stack) topIndex() int {
	return len(s.slice) - 1
}

type vars struct {
	objects map[string]Object
}

// Get retrieves a variable with the given identifier.
func (v *vars) Get(identifier string) (Object, error) {
	object, ok := v.objects[identifier]
	if !ok {
		return nil, NewVarNotFoundError(identifier)
	}
	return object, nil
}

// Set stores an object with the given identifier.
func (v *vars) Set(identifier string, object Object) error {
	v.objects[identifier] = object
	return nil
}

type VarNotFoundError struct{ Identifier string }

func NewVarNotFoundError(identifier string) VarNotFoundError {
	return VarNotFoundError{Identifier: identifier}
}

func (vnfe VarNotFoundError) Error() string {
	return fmt.Sprintf("variable '%v' not found", vnfe.Identifier)
}
