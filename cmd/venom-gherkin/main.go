package main

import (
	"log"

	"github.com/francois76/venom-gherkin/cmd/venom-gherkin/generate"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "venom-gherkin",
	Short: "venom-gherkin generate venom testsuites from gherkin files",
}

func main() {
	rootCmd.AddCommand(generate.Cmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Err:%s", err)
	}
}
