package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/angelofallars/hypo/internal/runtime"
)

const (
	splash = "Hypo interpreter (C) 2023"
	prompt = ">>> "
)

// Start starts the REPL environment.
func Start() {
	scanner := bufio.NewScanner(os.Stdin)
	runtime := runtime.New()

	fmt.Println(splash)
	for {
		fmt.Print(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		err := runtime.Eval(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
	}
}
