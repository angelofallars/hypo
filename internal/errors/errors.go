// package errs provides error types that occur in the runtime.
package errs

import "fmt"

type ErrorKind string

const (
	ParseKind     ErrorKind = "ParseError"
	StackKind     ErrorKind = "StackError"
	VariableKind  ErrorKind = "VariableError"
	TypeKind      ErrorKind = "TypeError"
	AttributeKind ErrorKind = "AttributeError"
)

// Dummy method
func (ek ErrorKind) errorKind() {}

type Error struct {
	message string
	kind    ErrorKind
}

func (e Error) Error() string {
	return fmt.Sprintf("%v: %v", e.kind, e.message)
}

func newHypoError(kind ErrorKind, message string, format []any) Error {
	return Error{
		message: fmt.Sprintf(message, format...),
		kind:    kind,
	}
}

// NewParseError returns a parsing error with a message.
func NewParseError(message string, format ...any) Error {
	return newHypoError(ParseKind, message, format)
}

// NewStackError returns a stack error with a message.
func NewStackError(message string, format ...any) Error {
	return newHypoError(StackKind, message, format)
}

// NewVariableError returns a variable error with a message.
func NewVariableError(message string, format ...any) Error {
	return newHypoError(VariableKind, message, format)
}

// NewTypeError returns a type error with a message.
func NewTypeError(message string, format ...any) Error {
	return newHypoError(TypeKind, message, format)
}

// NewAttributeError returns a type error with a message.
func NewAttributeError(message string, format ...any) Error {
	return newHypoError(AttributeKind, message, format)
}
