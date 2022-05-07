package transform

import (
	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformTestSteps(steps []*cucumber.Step) []venom.TestStep {
	result := make([]venom.TestStep, 0, len(steps))
	for _, step := range steps {
		result = append(result, TransformTestStep(step))
	}
	return result
}
