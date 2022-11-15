package command

import (
	"fmt"
	"task/db"

	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:   "list",
	Short: "list all task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something wrong list all task:", err)
		}

		for _, task := range tasks {
			fmt.Printf("%d.%s \n", task.Key, task.Value)
		}
	},
}
