package cmd

import (
	"fmt"
	"os"

	"github.com/angelofallars/hypo/internal/repl"
	"github.com/angelofallars/hypo/internal/runtime"
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

			err = runtime.New().Eval(contents)
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
