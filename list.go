package main

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all registered plugins",
	Run: func(cmd *cobra.Command, args []string) {
		names := make([]string, 0)
		for name := range registry {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			p := registry[name]
			name := p.Name()
			desc := p.Description()
			fmt.Printf("%s:\n\t%s\n", name, desc)
		}
	},
}
