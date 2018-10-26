package smartappcore_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSmartappcore(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Smartappcore Suite")
}
