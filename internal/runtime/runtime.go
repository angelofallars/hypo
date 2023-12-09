package runtime

import (
	"github.com/angelofallars/hypo/internal/evaluator"
	"github.com/angelofallars/hypo/internal/object"
	"github.com/angelofallars/hypo/internal/parser"
)

type Runtime struct {
	env *object.Env
}

func New() *Runtime {
	return &Runtime{
		env: object.NewEnv(),
	}
}

// Eval executes HTML, the programming language code from a string.
//
// State, like the stack and variable list, is maintained between Eval calls to the same [Runtime] instance.
func (i *Runtime) Eval(s string) error {
	program, err := parser.Parse(s)
	if err != nil {
		return err
	}

	err = evaluator.Exec(program, i.env)
	if err != nil {
		return err
	}

	return nil
}
