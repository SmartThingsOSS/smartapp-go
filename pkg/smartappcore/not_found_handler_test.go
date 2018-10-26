package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

var _ = Describe("NotFoundHandler", func() {

	var (
		handler *NotFoundHandler
	)

	BeforeEach(func() {
		handler = &NotFoundHandler{}
	})

	Context("With an instance of the default NotFoundHandler implementation", func() {
		It("should always Test positive", func() {
			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
				},
			})).To(Equal(true))

			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
				},
			})).To(Equal(true))
		})

		It("should return a 404", func() {
			response, err := handler.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(404)))
		})
	})
})
