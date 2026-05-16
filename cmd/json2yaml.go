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

		// ✅ validate file extension
		if err := validateExt(input, ".json"); err != nil {
			fmt.Println("Error:", err)
			return
		}

		// ✅ read file
		data, err := readFile(input)
		if err != nil {
			fmt.Println("File error:", err)
			return
		}

		// ✅ parse JSON
		var obj interface{}
		if err := json.Unmarshal(data, &obj); err != nil {
			fmt.Println("Invalid JSON format")
			return
		}

		// ✅ convert to YAML
		yamlData, err := yaml.Marshal(obj)
		if err != nil {
			fmt.Println("YAML conversion error:", err)
			return
		}

		// ⭐ PRINT OUTPUT (this is what you asked for)
		if printOutput {
			fmt.Println("----- YAML OUTPUT -----")
			fmt.Println(string(yamlData))
		}

		// ✅ decide output file name
		outFile := json2yamlOutput
		if outFile == "" {
			outFile = generateOutput(input, ".yaml")
		}

		// ✅ save file
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
