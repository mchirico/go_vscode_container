package idea

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe("Book", func() {
	Context("strings.Contains()", func() {
		When("the string contains the substring in the middle", func() {
			It("returns `true`", func() {
				Expect(strings.Contains("Ginkgo is awesome", "is")).To(BeTrue())
			})
		})
	})
})
