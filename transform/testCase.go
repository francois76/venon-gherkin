package transform

import (
	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformTestCase(child *cucumber.FeatureChild) venom.TestCase {
	result := venom.TestCase{
		Name:      child.Scenario.Name,
		TestSteps: TransformTestSteps(child.Scenario.Steps),
	}
	return result
}
