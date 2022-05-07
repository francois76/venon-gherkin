package generate

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Cmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate venom template",
	Long:    `venom-gherkin generate`,
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.LocalFlags().VisitAll(initFromCommandArguments)
		generateFiles()
	},
}

var (
	inputDir  string
	outputDir string

	inputDirFlag  *string
	outputDirFlag *string
)

func init() {
	inputDirFlag = Cmd.PersistentFlags().String("input-dir", "input", "Input Directory: taking gherkin files from this directory")
	outputDirFlag = Cmd.PersistentFlags().String("output-dir", "output", "Output Directory: writing venom files to this directory")
}
func initFromCommandArguments(f *pflag.Flag) {
	if !f.Changed {
		return
	}

	switch f.Name {
	case "input-dir":
		if inputDirFlag != nil {
			inputDir = *inputDirFlag
		}
	case "output-dir":
		if outputDirFlag != nil {
			outputDir = *outputDirFlag
		}
	}
}

func generateFiles() {
	fmt.Printf("Generate files from %s to %s", inputDir, outputDir)
}
