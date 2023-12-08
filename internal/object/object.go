package object

import "fmt"

type ObjectType string

const (
	NumberObject ObjectType = "Number"
	StringObject ObjectType = "String"
	BoolObject   ObjectType = "Bool"
	ArrayObject  ObjectType = "Array"
	ObjObject    ObjectType = "Obj"
	NullObject   ObjectType = "Null"
)

// Dummy method to make the type enum-like.
func (ot ObjectType) objectType() {}

// Object represents an object in the runtime.
type Object interface {
	// Type returns the primitive type of the object.
	Type() ObjectType
	// Display returns the string representation of the value.
	Display() string
}

type Number struct {
	Value float64
}

func (n *Number) Type() ObjectType { return NumberObject }
func (n *Number) Display() string  { return fmt.Sprint(n.Value) }

type String struct {
	Value string
}

func (n *String) Type() ObjectType { return StringObject }
func (n *String) Display() string  { return fmt.Sprintf("\"%v\"", n.Value) }

type Bool struct {
	Value bool
}

func (n *Bool) Type() ObjectType { return BoolObject }
func (n *Bool) Display() string  { return fmt.Sprint(n.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NullObject }
func (n *Null) Display() string  { return "null" }
