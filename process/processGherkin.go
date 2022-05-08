package process

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	gherkin_parser "github.com/cucumber/common/gherkin/go/v23"
)

func ProcessGherkinFile(inputDir string, outputDir string, fileName string) interface{} {
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
	return gherkinDocument

}
