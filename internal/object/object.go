package object

import (
	"fmt"
	"strings"

	"github.com/angelofallars/hypo/pkg/sliceutil"
)

type ObjectType string

const (
	NumberType ObjectType = "Number"
	StringType ObjectType = "String"
	BoolType   ObjectType = "Bool"
	ObjType    ObjectType = "Obj"
	NullType   ObjectType = "Null"
	ArrayType  ObjectType = "Array"
)

// Dummy method to make the type enum-like.
func (ot ObjectType) objectType() {}

// Object represents an object in the runtime.
type Object interface {
	// Type returns the primitive type of the object.
	Type() ObjectType
	// String returns the string representation of the value.
	String() string
}

type Number struct {
	Value float64
}

func (n *Number) Type() ObjectType { return NumberType }
func (n *Number) String() string   { return fmt.Sprint(n.Value) }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return StringType }
func (s *String) String() string   { return "\"" + s.Value + "\"" }

type Bool struct {
	Value bool
}

func (b *Bool) Type() ObjectType { return BoolType }
func (b *Bool) String() string   { return fmt.Sprint(b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NullType }
func (n *Null) String() string   { return "null" }

type Array struct {
	Value []Object
}

func (n *Array) Type() ObjectType { return ArrayType }
func (n *Array) String() string {
	displays := sliceutil.Map(n.Value,
		func(obj Object) string { return obj.String() },
	)

	return "[" + strings.Join(displays, ", ") + "]"
}
