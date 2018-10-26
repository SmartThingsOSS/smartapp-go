package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

type PingHandler struct {
}

func (h *PingHandler) Test(params *SmartAppParams) bool {
	request := params.Request
	return request != nil && request.Lifecycle == smartapp.AppLifecyclePING
}

func (h *PingHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	request := params.Request
	if request == nil || request.PingData == nil || *request.PingData.Challenge == "" {
		return nil, IllegalArgumentPingChallengeError
	}

	return &smartapp.ExecutionResponse{
		StatusCode: 200,
		PingData: &smartapp.PingResponseData{
			Challenge: *request.PingData.Challenge,
		},
	}, nil
}
