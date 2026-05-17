package cmd

import (
	"fmt"
	"vegeta/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all records from the database",
	Run: func(cmd *cobra.Command, args []string) {
		err := db.FetchAll()
		if err != nil {
			fmt.Println("Error fetching data:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
