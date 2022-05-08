package process

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type YamlTemplateObject struct {
	Templates []yamlTemplate ` yaml:"templates"`
}

type yamlTemplate struct {
	Given          string      ` yaml:"given"`
	When           string      ` yaml:"when"`
	Then           string      ` yaml:"then"`
	TemplateObject interface{} ` yaml:"template"`
}

func ProcessTemplate(inputDir string, outputDir string, fileName string) interface{} {
	content, err := ioutil.ReadFile(fmt.Sprint(inputDir, "/", fileName))
	if err != nil {
		log.Fatal(err)
	}
	var yamlTemplateObject YamlTemplateObject
	yaml.Unmarshal(content, &yamlTemplateObject)
	return yamlTemplateObject
}
