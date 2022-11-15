package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCommand = &cobra.Command{
	Use:   "do",
	Short: "do task",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("do task specified")
		fmt.Printf("%v", args)
		return nil
	},
}
