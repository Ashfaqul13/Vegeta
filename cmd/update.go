package cmd

import (
	"fmt"
	"os"
	"vegeta/db"

	"github.com/spf13/cobra"
)

var updateFilePath string

var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an existing record's content using a JSON file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		// Read the new file
		content, err := os.ReadFile(updateFilePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		// Update the DB
		err = db.UpdateRecord(id, string(content))
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("Successfully updated record ID %s\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateFilePath, "input", "i", "", "New JSON file for update")
	updateCmd.MarkFlagRequired("input")
}
