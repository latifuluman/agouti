package matchers

import (
	"github.com/onsi/gomega/types"
	"github.com/sclevine/agouti/matchers/internal/selection"
)

func HaveText(text string) types.GomegaMatcher {
	return &selection.HaveTextMatcher{ExpectedText: text}
}

func HaveAttribute(attribute string, value string) types.GomegaMatcher {
	return &selection.HaveAttributeMatcher{ExpectedAttribute: attribute, ExpectedValue: value}
}

func HaveCSS(property string, value string) types.GomegaMatcher {
	return &selection.HaveCSSMatcher{ExpectedProperty: property, ExpectedValue: value}
}

func BeSelected() types.GomegaMatcher {
	return &selection.BeSelectedMatcher{}
}

func BeVisible() types.GomegaMatcher {
	return &selection.BeVisibleMatcher{}
}
