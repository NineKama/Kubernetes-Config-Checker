package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-checker",
	Short: "Kubernetes Config Checker",
	Long:  `A tool for validating Kubernetes configuration files based on best practices.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `k8s-checker validate --file=<path>` to validate a YAML file.")
	},
}

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate Kubernetes YAML files",
	Long:  `This command validates Kubernetes YAML files to ensure they follow best practices.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")

		fmt.Println(filePath)
		if filePath == "" {
			fmt.Println("Error: --file flag is required.")
			cmd.Usage()
			os.Exit(1)
		}

		podSpec, err := loadYAMLFile(filePath)
		handleError(err, "Error reading or parsing YAML file")
		validateContainerResources(podSpec)
		validateSecurityContext(podSpec)
		validateServiceAccount(podSpec)
		validateNameSpace(podSpec)
		validateImmutable(podSpec)
	},
}

func init() {
	// Add the validate command to root
	rootCmd.AddCommand(validateCmd)

	// Add flags to validate command
	validateCmd.Flags().StringP("file", "f", "", "Path to the Kubernetes YAML file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func loadYAMLFile(filePath string) (podSpec, error) {
	var spec podSpec
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return spec, err
	}
	err = yaml.Unmarshal(yamlFile, &spec)
	return spec, err
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
