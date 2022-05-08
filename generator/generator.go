package generator

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin/transform"
	"gopkg.in/yaml.v3"
)

func GenerateGherkinDocument(fileName string, outputDir string, gherkinDocument *cucumber.GherkinDocument) {
	data, err := yaml.Marshal(transform.TransformTestSuite(gherkinDocument))
	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile(fmt.Sprint(outputDir, "/", strings.Replace(fileName, ".feature", ".venom.yaml", 1)), data, fs.FileMode(0755))
	if err2 != nil {
		log.Fatal(err2)
	}
}
