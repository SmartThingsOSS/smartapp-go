package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Authenticator", func() {
	Context("With an Authenticator", func() {
		It("should be able to create a new authenticator", func() {

			authenticator := NewAuthenticator(&AuthConfig{})

			Expect(authenticator).To(Not(BeNil()))
		})
	})
})
