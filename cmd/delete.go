package cmd

import (
	"fmt"
	"vegeta/db"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a record by ID",
	Args:  cobra.ExactArgs(1), // Ensures the user provides an ID
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		err := db.DeleteRecord(id)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Successfully deleted record with ID %s\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
