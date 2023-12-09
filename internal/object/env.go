package object

import (
	"errors"
	"slices"
)

// Env is an environment of the runtime which contains runtime values.
type Env struct {
	// Stack is the primary storage of values.
	Stack stack
	Vars  map[string]Object
}

type stack struct {
	slice []Object
}

// NewEnv returns a new [object.Env] instance.
func NewEnv() *Env {
	return &Env{
		Stack: stack{make([]Object, 0, 256)},
		Vars: map[string]Object{
			// Start with standard variables for common values
			"true":  &Bool{Value: true},
			"false": &Bool{Value: false},
			"null":  &Null{},
		},
	}
}

var ErrStackEmpty = errors.New("stack empty")

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
