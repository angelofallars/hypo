package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/angelofallars/hypo/internal/evaluator"
	"github.com/angelofallars/hypo/internal/object"
	"github.com/angelofallars/hypo/internal/parser"
)

const (
	splash = "Hypo interpreter (C) 2023"
	prompt = ">>> "
)

// Start starts the REPL environment.
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	env := object.NewEnv()

	fmt.Println(splash)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		node, err := parser.Parse(line)
		if err != nil {
			fmt.Errorf("%w\n", err)
			continue
		}

		err = evaluator.Exec(node, env)
		if err != nil {
			fmt.Errorf("%w\n", err)
			continue
		}
	}
}
