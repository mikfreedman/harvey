package operation_test

import (
	"testing"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

var suite spec.Suite

func init() {
	suite = spec.New("operation", spec.Report(report.Terminal{}))
	suite("search", testStripFrontMatter)
}

func TestOperation(t *testing.T) {
	suite.Run(t)
}
