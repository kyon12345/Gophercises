package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCommand = &cobra.Command{
	Use:     "compeleted",
	Aliases: []string{"cmp"},
	Short:   "list all completed",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("list all completed task")
		return nil
	},
}