package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"net"
)

type PredicateEventHandler interface {
	Test(event *smartapp.Event) bool
	Handle(params *SmartAppParams, event *smartapp.Event) error
}

type OnSuccess func(params *SmartAppParams) (*smartapp.ExecutionResponse, error)
type OnError func(params *SmartAppParams, err error) (*smartapp.ExecutionResponse, error)
type FailOnError func(err error) bool

type EventHandler interface {
	SetHandlerChain(chain []PredicateEventHandler)
	AddEventHandler(handler PredicateEventHandler)
	DoOnSuccess(onSuccess OnSuccess)
	DoOnError(onError OnError)
	SetFailOnErrorPredicate(failOnError FailOnError)
	GetChainSize() int
	BuildHandler() Handler
}

type DefaultEventHandler struct {
	Chain       []PredicateEventHandler
	OnSuccess   OnSuccess
	OnError     OnError
	FailOnError FailOnError
}

func NewEventHandler() EventHandler {
	return &DefaultEventHandler{
		Chain: make([]PredicateEventHandler, 0, 0),
		OnSuccess: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return &smartapp.ExecutionResponse{
				StatusCode: 200,
				EventData:  net.Interface{},
			}, nil
		},
		OnError: func(params *SmartAppParams, err error) (*smartapp.ExecutionResponse, error) {
			return &smartapp.ExecutionResponse{
				StatusCode: 500,
			}, err
		},
		FailOnError: func(err error) bool {
			return true
		},
	}
}

func (h *DefaultEventHandler) SetHandlerChain(chain []PredicateEventHandler) {
	newChain := make([]PredicateEventHandler, 0, 0)
	if chain != nil && len(chain) > 0 {
		newChain = append(newChain, chain...)
	}
	h.Chain = newChain
}

func (h *DefaultEventHandler) AddEventHandler(handler PredicateEventHandler) {
	if handler != nil {
		h.Chain = append(h.Chain, handler)
	}
}
func (h *DefaultEventHandler) DoOnSuccess(onSuccess OnSuccess) {
	if onSuccess != nil {
		h.OnSuccess = onSuccess
	}
}

func (h *DefaultEventHandler) DoOnError(onError OnError) {
	if onError != nil {
		h.OnError = onError
	}
}

func (h *DefaultEventHandler) SetFailOnErrorPredicate(failOnError FailOnError) {
	if failOnError != nil {
		h.FailOnError = failOnError
	}
}

func (h *DefaultEventHandler) Test(params *SmartAppParams) bool {
	request := params.Request
	return request != nil && request.Lifecycle == smartapp.AppLifecycleEVENT
}

func (h *DefaultEventHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	request := params.Request

	if request == nil || request.Lifecycle != smartapp.AppLifecycleEVENT {
		return nil, UnsupportedLifecyleError
	}

	if request.EventData == nil || len(request.EventData.Events) == 0 {
		return &smartapp.ExecutionResponse{
			StatusCode: 200,
		}, nil
	}

	for _, event := range request.EventData.Events {
		handler := h.findEventHandler(event)
		if handler != nil {
			if err := handler.Handle(params, event); err != nil {
				if h.FailOnError(err) {
					return h.OnError(params, err)
				}
			}
		}
	}

	return h.OnSuccess(params)
}

func (h *DefaultEventHandler) BuildHandler() Handler {
	return func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
		return h.Handle(params)
	}
}

func (h *DefaultEventHandler) GetChainSize() int {
	return len(h.Chain)
}

func (h *DefaultEventHandler) findEventHandler(event *smartapp.Event) PredicateEventHandler {
	if h.Chain != nil || len(h.Chain) > 0 {
		for _, handler := range h.Chain {
			if handler.Test(event) {
				return handler
			}
		}
	}
	return nil
}
