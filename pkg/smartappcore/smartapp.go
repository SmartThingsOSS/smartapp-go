package smartappcore

import (
	"context"
	"errors"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"log"
)

type DefaultSmartApp struct {
	Chain []PredicateHandler
}

const (
	CorrelationHeader = "X-ST-CORRELATION"
	CorrelationKey    = "loggingId"
	MediaTypeJsonV1   = "application/vnd.smartthings+json;v=1"
	SmartThingsHost   = "api.smartthings.com"
)

var (
	UnsupportedLifecyleError          = errors.New("unsupported_app_lifecycle")
	IllegalArgumentPingChallengeError = errors.New("illegal_argument_ping_challenge")
)

type SmartAppParams struct {
	Request *smartapp.ExecutionRequest
	Context context.Context
}

func NewSmartApp(definition SmartAppDefinition) SmartApp {
	if definition == nil {
		log.Panic("Cannot launch SmartApp with a definition!")
	}

	return &DefaultSmartApp{
		Chain: definition.BuildChain(),
	}
}

func (p *SmartAppParams) GetSetting(key string) string {
	if p.Request != nil && p.Request.Settings != nil {
		return p.Request.Settings[key]
	}
	return ""
}

type SmartApp interface {
	Handle(request *SmartAppParams) (*smartapp.ExecutionResponse, error)
}

func (s *DefaultSmartApp) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	for _, handler := range s.Chain {
		if handler.Test(params) {
			return handler.Handle(params)
		}
	}
	return nil, UnsupportedLifecyleError
}
