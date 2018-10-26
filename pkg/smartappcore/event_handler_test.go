package smartappcore

import (
	"errors"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type TestEventHandler struct {
	Count int
	Value bool
	Error error
}

func (h *TestEventHandler) Test(event *smartapp.Event) bool {
	return h.Value
}

func (h *TestEventHandler) Handle(params *SmartAppParams, event *smartapp.Event) error {
	h.Count = h.Count + 1
	return h.Error
}

var _ = Describe("PredicateHandler", func() {
	var (
		handler EventHandler
	)

	BeforeEach(func() {
		handler = NewEventHandler()
	})

	Context("With an instance of the default EventHandler implementation", func() {
		It("should have 0 handlers in it's chain by default", func() {
			Expect(handler.GetChainSize()).To(Equal(0))
		})

		It("should allow me to set entire handler chain", func() {
			chain := make([]PredicateEventHandler, 0, 3)
			chain = append(chain, &TestEventHandler{
				Value: false,
			})
			chain = append(chain, &TestEventHandler{
				Value: false,
			})
			chain = append(chain, &TestEventHandler{
				Value: false,
			})

			Expect(handler.GetChainSize()).To(Equal(0))
			handler.SetHandlerChain(chain)
			Expect(handler.GetChainSize()).To(Equal(3))
		})

		It("should allow me to add a singlehandler", func() {
			Expect(handler.GetChainSize()).To(Equal(0))
			handler.AddEventHandler(&TestEventHandler{
				Value: false,
			})
			Expect(handler.GetChainSize()).To(Equal(1))
			handler.AddEventHandler(&TestEventHandler{
				Value: false,
			})
			Expect(handler.GetChainSize()).To(Equal(2))
			handler.AddEventHandler(&TestEventHandler{
				Value: false,
			})
			Expect(handler.GetChainSize()).To(Equal(3))
		})

		It("should call a success handler even when no predicates matched", func() {
			success := false
			h := &TestEventHandler{
				Value: false,
			}
			handler.AddEventHandler(h)
			handler.DoOnSuccess(func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
				success = true
				return &smartapp.ExecutionResponse{StatusCode: 200}, nil
			})
			Expect(handler.GetChainSize()).To(Equal(1))
			Expect(success).To(Equal(false))
			Expect(h.Count).To(Equal(0))

			handlerFunction := handler.BuildHandler()
			response, err := handlerFunction(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
					EventData: &smartapp.EventData{
						Events: []*smartapp.Event{
							{
								EventType: smartapp.EventTypeMODEEVENT,
							},
						},
					},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(success).To(Equal(true))
			Expect(h.Count).To(Equal(0))
		})

		It("should skip processing and return 200 if no event sent on request", func() {
			success := false
			h := &TestEventHandler{
				Value: false,
			}
			handler.AddEventHandler(h)
			handler.DoOnSuccess(func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
				success = true
				return &smartapp.ExecutionResponse{StatusCode: 500}, nil
			})
			Expect(handler.GetChainSize()).To(Equal(1))
			Expect(success).To(Equal(false))
			Expect(h.Count).To(Equal(0))

			handlerFunction := handler.BuildHandler()
			response, err := handlerFunction(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(success).To(Equal(false))
			Expect(h.Count).To(Equal(0))
		})

		It("should invoke event handler when predicate matched", func() {
			success := false
			h := &TestEventHandler{
				Value: true,
			}
			handler.AddEventHandler(h)
			handler.DoOnSuccess(func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
				success = true
				return &smartapp.ExecutionResponse{StatusCode: 200}, nil
			})
			Expect(handler.GetChainSize()).To(Equal(1))
			Expect(success).To(Equal(false))
			Expect(h.Count).To(Equal(0))

			handlerFunction := handler.BuildHandler()
			response, err := handlerFunction(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
					EventData: &smartapp.EventData{
						Events: []*smartapp.Event{
							{
								EventType: smartapp.EventTypeMODEEVENT,
							},
						},
					},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(success).To(Equal(true))
			Expect(h.Count).To(Equal(1))
		})
	})

	It("should invoke on error handler", func() {
		success := false
		failure := false
		myError := errors.New("an_error")
		h := &TestEventHandler{
			Value: true,
			Error: myError,
		}
		handler.AddEventHandler(h)
		handler.DoOnSuccess(func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			success = true
			return &smartapp.ExecutionResponse{StatusCode: 200}, nil
		})
		handler.DoOnError(func(params *SmartAppParams, err error) (*smartapp.ExecutionResponse, error) {
			failure = true
			return &smartapp.ExecutionResponse{StatusCode: 500}, nil
		})
		Expect(handler.GetChainSize()).To(Equal(1))
		Expect(success).To(Equal(false))
		Expect(h.Count).To(Equal(0))

		handlerFunction := handler.BuildHandler()
		response, err := handlerFunction(&SmartAppParams{
			Request: &smartapp.ExecutionRequest{
				Lifecycle: smartapp.AppLifecycleEVENT,
				EventData: &smartapp.EventData{
					Events: []*smartapp.Event{
						{
							EventType: smartapp.EventTypeMODEEVENT,
						},
					},
				},
			},
		})

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(int64(500)))
		Expect(success).To(Equal(false))
		Expect(failure).To(Equal(true))
		Expect(h.Count).To(Equal(1))
	})

	It("should support multiple event handlers", func() {
		success := false
		failure := false
		h1 := &TestEventHandler{
			Value: false,
		}
		h2 := &TestEventHandler{
			Value: true,
		}

		handler.AddEventHandler(h1)
		handler.AddEventHandler(h2)
		handler.DoOnSuccess(func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			success = true
			return &smartapp.ExecutionResponse{StatusCode: 201}, nil
		})
		handler.DoOnError(func(params *SmartAppParams, err error) (*smartapp.ExecutionResponse, error) {
			failure = true
			return &smartapp.ExecutionResponse{StatusCode: 500}, nil
		})
		Expect(handler.GetChainSize()).To(Equal(2))
		Expect(success).To(Equal(false))
		Expect(failure).To(Equal(false))
		Expect(h1.Count).To(Equal(0))
		Expect(h2.Count).To(Equal(0))

		handlerFunction := handler.BuildHandler()
		response, err := handlerFunction(&SmartAppParams{
			Request: &smartapp.ExecutionRequest{
				Lifecycle: smartapp.AppLifecycleEVENT,
				EventData: &smartapp.EventData{
					Events: []*smartapp.Event{
						{
							EventType: smartapp.EventTypeMODEEVENT,
						},
					},
				},
			},
		})

		Expect(err).To(BeNil())
		Expect(response.StatusCode).To(Equal(int64(201)))
		Expect(success).To(Equal(true))
		Expect(failure).To(Equal(false))
		Expect(h1.Count).To(Equal(0))
		Expect(h2.Count).To(Equal(1))
	})
})
