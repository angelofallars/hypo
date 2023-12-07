package main

import (
	"os"

	"github.com/angelofallars/hypo/internal/cmd"
)

func main() {
	statusCode := cmd.Exec()
	os.Exit(statusCode)
}
