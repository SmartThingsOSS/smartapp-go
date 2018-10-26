package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

var _ = Describe("PingHandler", func() {

	var (
		handler *PingHandler
	)

	BeforeEach(func() {
		handler = &PingHandler{}
	})

	Context("With an instance of the default PingHandler implementation", func() {
		It("should Test positive only for PING lifecycle", func() {
			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
				},
			})).To(Equal(true))

			Expect(handler.Test(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
				},
			})).To(Equal(false))
		})

		It("should successfully return challenge", func() {
			aChallenge := "challenge"
			response, err := handler.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
					PingData: &smartapp.PingData{
						Challenge: &aChallenge,
					},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(response.PingData.Challenge).To(Equal(aChallenge))
		})
	})
})
