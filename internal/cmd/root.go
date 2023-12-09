package cmd

import (
	"os"

	"github.com/angelofallars/hypo/internal/repl"
	"github.com/angelofallars/hypo/internal/runtime"
	"github.com/spf13/cobra"
)

func Exec() int {
	rootCmd := &cobra.Command{
		Use:          "hypo [ file ]",
		Short:        "Hypo is a fast runtime for HTML, the programming language running outside the browser.",
		Args:         cobra.MaximumNArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				repl.Start()
				return nil
			}

			bytes, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			contents := string(bytes)

			err = runtime.New().Eval(contents)
			if err != nil {
				return err
			}

			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
