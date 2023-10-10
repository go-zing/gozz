package main

import (
	"os"

	"github.com/spf13/cobra"

	_ "github.com/Just-maple/gozz/internal/plugins"
	"github.com/Just-maple/gozz/zcore"
)

var (
	cmd = cobra.Command{
		Use:          zcore.ExecName,
		SilenceUsage: true,
	}

	registry = zcore.PluginRegistry()
)

func init() {
	cmd.AddCommand(
		run,
		listCmd,
	)
}

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
