package cmd

import (
	"fmt"
	"os"

	"github.com/Aristat/go-testing/cmd/file_examples"

	"github.com/Aristat/go-testing/cmd/chan_examples"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "bin [command]",
	Long:          "",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Example codess",
}

func Execute() {
	rootCmd.AddCommand(chan_examples.Cmd, file_examples.Cmd)
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}
