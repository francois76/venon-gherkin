package generate

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	gherkin_parser "github.com/cucumber/common/gherkin/go/v23"
	"github.com/francois76/venom-gherkin/transform"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
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
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		content, err := ioutil.ReadFile(fmt.Sprint(inputDir, "/", file.Name()))
		if err != nil {
			log.Fatal(err)
		}
		// convert byte slice to io.Reader
		reader := bytes.NewReader(content)
		gherkinDocument, err := gherkin_parser.ParseGherkinDocument(reader, func() string {
			return "1"
		})
		if err != nil {
			log.Fatal(err)
		}
		data, err := yaml.Marshal(transform.TransformFeature(gherkinDocument))
		if err != nil {
			log.Fatal(err)
		}

		err2 := ioutil.WriteFile(fmt.Sprint(outputDir, "/", file.Name()), data, 0)

		if err2 != nil {

			log.Fatal(err2)
		}
	}
}
