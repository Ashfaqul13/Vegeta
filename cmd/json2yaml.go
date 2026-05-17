package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var json2yamlOutput string
var printOutput bool

var json2yamlCmd = &cobra.Command{
	Use:   "json2yaml [file]",
	Short: "Convert JSON to YAML",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		input := args[0]

		if err := validateExt(input, ".json"); err != nil {
			fmt.Println("Error:", err)
			return
		}

		data, err := readFile(input)
		if err != nil {
			fmt.Println("File error:", err)
			return
		}

		var obj interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			fmt.Println("Invalid JSON format")
			return
		}

		yamlData, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Println("YAML conversion error:", err)
			return
		}

		if printOutput {
			fmt.Println("----- YAML OUTPUT -----")
			fmt.Println(string(yamlData))
		}

		outFile := json2yamlOutput
		if outFile == "" {

		}

		err = os.WriteFile(outFile, yamlData, 0644)
		if err != nil {
			fmt.Println("Write file error:", err)
			return
		}

		fmt.Println("Converted to:", outFile)
	},
}

func init() {
	json2yamlCmd.Flags().StringVarP(&json2yamlOutput, "output", "o", "", "output file name")
	json2yamlCmd.Flags().BoolVarP(&printOutput, "print", "p", false, "print output in terminal")

	rootCmd.AddCommand(json2yamlCmd)
}
