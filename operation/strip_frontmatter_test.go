package operation_test

import (
	"bytes"
	"testing"

	"github.com/mikfreedman/harvey/operation"
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
)

func testStripFrontMatter(t *testing.T, when spec.G, it spec.S) {

	var input *bytes.Buffer
	var output *bytes.Buffer
	it.Before(func() {
		RegisterTestingT(t)
		output = new(bytes.Buffer)
	})

	when("the input does not contain frontmatter", func() {
		it.Before(func() {
			input = bytes.NewBufferString("this is just a test\nline two\n")
		})

		it("returns an error", func() {
			operation.StripFrontMatter(input, output)
			Expect(output.String()).To(Equal("this is just a test\nline two\n"))
		})
	})

	when("the does contain frontmatter", func() {
		it.Before(func() {
			input = bytes.NewBufferString("---\nthis: is frontmatter\n---\n\nthis is just a test\nline two\n")
		})

		it("returns an error", func() {
			operation.StripFrontMatter(input, output)
			Expect(output.String()).To(Equal("this is just a test\nline two\n"))
		})
	})
}
