package process

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type YamlTemplateObject struct {
	templates []yamlTemplate ` yaml:"templates"`
}

type yamlTemplate struct {
	given          string      ` yaml:"given"`
	when           string      ` yaml:"when"`
	then           string      ` yaml:"then"`
	templateObject interface{} ` yaml:"template"`
}

func ProcessTemplate(inputDir string, outputDir string, fileName string) {
	content, err := ioutil.ReadFile(fmt.Sprint(inputDir, "/", fileName))
	if err != nil {
		log.Fatal(err)
	}
	var yamlTemplateObject YamlTemplateObject
	yaml.Unmarshal(content, &yamlTemplateObject)
}
