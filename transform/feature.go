package transform

import (
	messages "github.com/cucumber/common/messages/go/v18"
	"github.com/francois76/venom-gherkin"
)

func TransformFeature(feature *messages.GherkinDocument) *venom.TestSuite {
	result := &venom.TestSuite{
		Name: feature.Feature.Name,
	}
	return result
}
