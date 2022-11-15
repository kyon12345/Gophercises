package command

import (
	"fmt"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

var rmCommand = &cobra.Command{
	Use:   "rm",
	Short: "remove one task from list",
	Run: func(cmd *cobra.Command, args []string) {
		key, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		if err := db.DeleteTask(key); err != nil {
			fmt.Println("error while delete:", err)
		}
	},
}
