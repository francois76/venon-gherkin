package generate

import (
	"fmt"
	"io/ioutil"
	"log"

	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin/generator"
	"github.com/francois76/venom-gherkin/process"
	"github.com/francois76/venom-gherkin/template"
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
	fmt.Printf("Generate files from %s to %s\n", inputDir, outputDir)

	fileProcessorFunc := func(in string, singleProcessor func(inputDir string, outputDir string, fileName string) interface{}) map[string]interface{} {
		files, err := ioutil.ReadDir(in)
		if err != nil {
			log.Fatal(err)
		}
		processedFiles := map[string]interface{}{}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			processedFile := singleProcessor(in, outputDir, file.Name())
			processedFiles[file.Name()] = processedFile
		}
		return processedFiles
	}

	// processing templates
	templates := fileProcessorFunc(fmt.Sprint(inputDir, "/", "templates"), process.ProcessTemplate)
	template.RegisterTemplates(templates)
	// processing gherkin files
	gherkinDocuments := fileProcessorFunc(inputDir, process.ProcessGherkinFile)
	for fileName, gherkinDocument := range gherkinDocuments {
		generator.GenerateGherkinDocument(fileName, outputDir, gherkinDocument.(*cucumber.GherkinDocument))
	}
}
