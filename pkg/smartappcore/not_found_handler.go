package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

type NotFoundHandler struct {
}

func (h *NotFoundHandler) Test(params *SmartAppParams) bool {
	return true
}

func (h *NotFoundHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	return &smartapp.ExecutionResponse{
		StatusCode: 404,
	}, nil
}
