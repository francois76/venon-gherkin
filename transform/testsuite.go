package transform

import (
	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformTestSuite(feature *cucumber.GherkinDocument) *venom.TestSuite {
	result := &venom.TestSuite{
		Name:      feature.Feature.Name,
		TestCases: TransformTestCases(feature.Feature.Children),
	}
	return result
}
