package idea_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIdea(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Idea Suite")
}
