package cmd

import (
	"fmt"
	"os"
	"vegeta/db" // Adjust based on your module name in go.mod

	"github.com/spf13/cobra"
)

var inputFilePath string

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves JSON data to Postgres",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Read file
		content, err := os.ReadFile(inputFilePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// 2. Connect to DB
		database, err := db.ConnectDB()
		if err != nil {
			fmt.Println("Could not connect to database. Is it running? Error:", err)
			return // This RETURN prevents the panic!
		}
		defer database.Close()

		// 3. Insert Data
		query := `INSERT INTO resource_table (content) VALUES ($1)`
		_, err = database.Exec(query, string(content))
		if err != nil {
			fmt.Println("Insert failed:", err)
			return
		}

		fmt.Println("--- SUCCESS ---")
		fmt.Println("Data dumped into resource_table.")
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().StringVarP(&inputFilePath, "input", "i", "", "JSON file to save")
	saveCmd.MarkFlagRequired("input")
}
