package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

var _ = Describe("NotFoundHandler", func() {

	var (
		handler *NoopLifecycleHandler
	)

	BeforeEach(func() {
		handler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleCONFIGURATION,
		}
	})

	Context("With an instance of the default NoopLifecycleHandler implementation", func() {
		It("should always Test positive for it's configured lifecycle", func() {
			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleCONFIGURATION,
				},
			})).To(Equal(true))

			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
				},
			})).To(Equal(false))
		})

		It("should always return a 200", func() {
			response, err := handler.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleCONFIGURATION,
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
		})
	})
})
