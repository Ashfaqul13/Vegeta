package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var yaml2jsonOutput string
var yamlPrintOutput bool

var yaml2jsonCmd = &cobra.Command{
	Use:   "yaml2json [file]",
	Short: "Convert YAML to JSON",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		if err := validateExt(input, ".yaml", ".yml"); err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := readFile(input)
		if err != nil {
			fmt.Println("File error:", err)
			return
		}

		var obj interface{}
		if err := yaml.Unmarshal(data, &obj); err != nil {
			fmt.Println("Invalid YAML format")
			return
		}

		jsonData, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			fmt.Println("JSON conversion error:", err)
			return
		}

		if yamlPrintOutput {
			fmt.Println("----- JSON OUTPUT -----")
			fmt.Println(string(jsonData))
		}

		outFile := yaml2jsonOutput
		if outFile == "" {

		}

		err = os.WriteFile(outFile, jsonData, 0644)
		if err != nil {
			fmt.Println("Write file error:", err)
			return
		}

		fmt.Println("Converted to:", outFile)
	},
}

func init() {
	yaml2jsonCmd.Flags().StringVarP(&yaml2jsonOutput, "output", "o", "", "output file name")
	yaml2jsonCmd.Flags().BoolVarP(&yamlPrintOutput, "print", "p", false, "print output in terminal")

	rootCmd.AddCommand(yaml2jsonCmd)
}
