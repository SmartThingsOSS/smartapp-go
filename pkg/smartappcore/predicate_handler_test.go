package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

var _ = Describe("PredicateHandler", func() {

	var (
		handler PredicateHandler
	)

	BeforeEach(func() {
		handler = &DefaultPredicateHandler{
			Tester: func(params *SmartAppParams) bool {
				return params.Request.Lifecycle == smartapp.AppLifecyclePING
			},
			Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
				return &smartapp.ExecutionResponse{
					StatusCode: 200,
				}, nil
			},
		}
	})

	Context("With an instance of the default PredicateHandler implementation", func() {
		It("should successfully evaluate truthy when test passes", func() {
			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
				},
			})).To(Equal(true))
		})

		It("should successfully evaluate false when test fails", func() {
			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleINSTALL,
				},
			})).To(Equal(false))
		})

		It("should support execution of handler", func() {
			response, err := handler.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleINSTALL,
				},
			})
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
		})
	})
})
