package template

import "github.com/francois76/venom-gherkin/process"

type TemplateItem struct {
	Expression     string
	Variables      map[string]string
	TemplateObject interface{}
}

type TemplatesStruct struct {
	Given []*TemplateItem
	When  []*TemplateItem
	Then  []*TemplateItem
}

var Templates = TemplatesStruct{
	Given: []*TemplateItem{},
	When:  []*TemplateItem{},
	Then:  []*TemplateItem{},
}

func RegisterTemplates(templates map[string]interface{}) {
	for _, templateItem := range templates {
		registerTemplateItem(templateItem.(process.YamlTemplateObject))
	}
}

func registerTemplateItem(item process.YamlTemplateObject) {
	for _, template := range item.Templates {
		if template.Given != "" {

		} else if template.When != "" {

		} else if template.Then != "" {

		}
	}
}
