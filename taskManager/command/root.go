package command

import "github.com/spf13/cobra"

var root = cobra.Command{
	Use:   "task",
	Short: "task is a CLI for managing your TODOs.",
}

func init() {
	root.AddCommand(addCommand)
	root.AddCommand(listCommand)
	root.AddCommand(doCommand)
	root.AddCommand(completeCommand)
	root.AddCommand(rmCommand)
}

func Execute() error {
	return root.Execute()
}
