package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "workspace",
	Short: "workspace tool to manage ccd git repositories",
}

func Execute() error {
	return RootCmd.Execute()
}
