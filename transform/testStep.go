package transform

import (
	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformTestStep(step *cucumber.Step) venom.TestStep {
	result := venom.TestStep{}
	result["what"] = step.Text
	return result
}
