package generate

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate venom template",
	Long:    `venom-gherkin generate`,
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		generateFiles()
	},
}

func generateFiles() {
	fmt.Println("Generate files")
}
