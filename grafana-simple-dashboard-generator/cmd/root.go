package root

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "todo",
  Short: "TODO short",
  Long: `TODO long`,
  Run: func(cmd *cobra.Command, args []string) {
    // TODO: Do Stuff Here
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
