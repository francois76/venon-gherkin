package transform

import (
	cucumber "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformTestCases(children []*cucumber.FeatureChild) []venom.TestCase {
	result := make([]venom.TestCase, 0, len(children))
	for _, child := range children {
		result = append(result, TransformTestCase(child))
	}
	return result
}
