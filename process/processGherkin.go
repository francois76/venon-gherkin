package process

import (
	"bytes"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	gherkin_parser "github.com/cucumber/common/gherkin/go/v23"
	"github.com/francois76/venom-gherkin/transform"
	"gopkg.in/yaml.v3"
)

func ProcessGherkinFile(inputDir string, outputDir string, fileName string) {
	content, err := ioutil.ReadFile(fmt.Sprint(inputDir, "/", fileName))
	if err != nil {
		log.Fatal(err)
	}
	reader := bytes.NewReader(content)
	gherkinDocument, err := gherkin_parser.ParseGherkinDocument(reader, func() string {
		return "1"
	})
	if err != nil {
		log.Fatal(err)
	}
	data, err := yaml.Marshal(transform.TransformTestSuite(gherkinDocument))
	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile(fmt.Sprint(outputDir, "/", strings.Replace(fileName, ".feature", ".venom.yaml", 1)), data, fs.FileMode(0755))
	if err2 != nil {
		log.Fatal(err2)
	}
}
