package cmd

import (
	"fmt"
	"os"

	"github.com/angelofallars/hypo/internal/evaluator"
	"github.com/angelofallars/hypo/internal/object"
	"github.com/angelofallars/hypo/internal/parser"
	"github.com/angelofallars/hypo/internal/repl"
	"github.com/spf13/cobra"
)

func Exec() int {
	rootCmd := &cobra.Command{
		Use:   "hypo [ file ]",
		Short: "Hypo is a fast runtime for HTML, the programming language running outside the browser.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				repl.Start()
			}

			bytes, err := os.ReadFile(args[0])
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			contents := string(bytes)

			node, err := parser.Parse(contents)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}

			env := object.NewEnv()

			err = evaluator.Exec(node, env)
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}

	return 0
}
