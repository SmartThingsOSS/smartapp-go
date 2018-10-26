package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

type Test func(params *SmartAppParams) bool
type Handle func(params *SmartAppParams) (*smartapp.ExecutionResponse, error)

type PredicateHandler interface {
	Test(params *SmartAppParams) bool
	Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error)
}

type DefaultPredicateHandler struct {
	Tester  Test
	Handler Handle
}

func (h *DefaultPredicateHandler) Test(params *SmartAppParams) bool {
	return h.Tester(params)
}

func (h *DefaultPredicateHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	return h.Handler(params)
}
