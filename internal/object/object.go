package object

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
	// Dummy method
	object()
	Type() ObjectType
}

type Number struct {
	Value float64
}

func (n *Number) object() {}
func (n *Number) Type() ObjectType {
	return NumberObject
}

type String struct {
	Value string
}

func (n *String) object() {}
func (n *String) Type() ObjectType {
	return StringObject
}

type Bool struct {
	Value bool
}

func (n *Bool) object() {}
func (n *Bool) Type() ObjectType {
	return BoolObject
}
