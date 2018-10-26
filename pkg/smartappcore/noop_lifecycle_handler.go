package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

type NoopLifecycleHandler struct {
	Lifecycle smartapp.AppLifecycle
}

func (h *NoopLifecycleHandler) Test(params *SmartAppParams) bool {
	request := params.Request
	return request != nil && request.Lifecycle == h.Lifecycle
}

func (h *NoopLifecycleHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	return &smartapp.ExecutionResponse{
		StatusCode: 200,
	}, nil
}
