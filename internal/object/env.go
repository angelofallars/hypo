package object

import (
	"errors"
	"slices"
)

// TODO: replace any with Object

// Env is an environment of the runtime which contains runtime values.
type Env struct {
	// Stack is the primary storage of values.
	Stack stack
	Vars  map[string]any
}

type stack struct {
	slice []any
}

func NewEnv() *Env {
	return &Env{
		Stack: stack{make([]any, 0, 256)},
		Vars: map[string]any{
			// Start with standard variables for common values
			"true":  true,
			"false": false,
			"null":  nil,
		},
	}
}

var ErrStackEmpty = errors.New("Stack empty")

// Push a value to the top of the stack.
func (s *stack) Push(v any) {
	s.slice = append(s.slice, v)
}

// Pop a value from the top of the stack and return it.
func (s *stack) Pop() (any, error) {
	v, err := s.Top()
	if err != nil {
		return nil, err
	}

	topIndex := s.topIndex()
	s.slice = slices.Delete(s.slice, topIndex, topIndex+1)
	return v, nil
}

// Receive a value from the top of the stack without consuming it.
func (s *stack) Top() (any, error) {
	if len(s.slice) == 0 {
		return nil, ErrStackEmpty
	}

	return s.slice[s.topIndex()], nil
}

// topIndex returns the last index in the stack.
func (s *stack) topIndex() int {
	return len(s.slice) - 1
}
