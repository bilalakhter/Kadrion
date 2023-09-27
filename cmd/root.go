/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmd = &cobra.Command{
	Use: "kadrion",
	Long: ` kadrion is a tool for testops

 Usage:

    kadrion <command> [arguments]

 The commands are:

    help         get help <argument> for command help
    performance  test performance of api

 Kindly visit kadriontestops.tech/docs for detailed documentation.`,
}

var processCmd = &cobra.Command{
	Use:   "process [yamlFile]",
	Short: "Process a YAML file",
	Args:  cobra.ExactArgs(1), // Require exactly 1 argument (the YAML file)
	Run: func(cmd *cobra.Command, args []string) {
		yamlFile := args[0]
		if err := processYAMLFile(yamlFile); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	},
}

func processYAMLFile(yamlFile string) error {
	// Open and read the YAML file
	file, err := os.Open(yamlFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Define a struct to unmarshal the YAML into
	var data interface{}

	// Create a YAML decoder and unmarshal the file into the struct
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	// Process the data (for now, print it)
	fmt.Printf(" %+v\n", data)

	return nil
}

func Execute() {
	rootCmd.AddCommand(processCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
