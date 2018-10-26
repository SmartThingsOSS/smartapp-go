package smartappcore

import (
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"log"
)

type Handler func(params *SmartAppParams) (*smartapp.ExecutionResponse, error)

type SmartAppDefinition interface {
	SetInstallHandler(handler Handler)
	SetUpdateHandler(handler Handler)
	SetUninstallHandler(handler Handler)
	SetEventHandler(handler Handler)
	SetPingHandler(handler Handler)
	SetConfigurationHandler(handler Handler)
	SetOAuthCallbackHandler(handler Handler)
	AddHandler(handler PredicateHandler)
	BuildChain() []PredicateHandler
}

type DefaultSmartAppDefinition struct {
	InstallHandler       PredicateHandler
	UpdateHandler        PredicateHandler
	UninstallHandler     PredicateHandler
	EventHandler         PredicateHandler
	PingHandler          PredicateHandler
	ConfigurationHandler PredicateHandler
	OAuthCallbackHandler PredicateHandler
	AdditionalHandlers   []PredicateHandler
}

func NewSmartAppDefinition() SmartAppDefinition {
	return &DefaultSmartAppDefinition{}
}

func (h *DefaultSmartAppDefinition) SetInstallHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleINSTALL
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.InstallHandler = handle
}

func (h *DefaultSmartAppDefinition) SetUpdateHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleUPDATE
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.UpdateHandler = handle
}

func (h *DefaultSmartAppDefinition) SetUninstallHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleUNINSTALL
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.UninstallHandler = handle
}

func (h *DefaultSmartAppDefinition) SetEventHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleEVENT
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.EventHandler = handle
}

func (h *DefaultSmartAppDefinition) SetPingHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecyclePING
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.PingHandler = handle
}

func (h *DefaultSmartAppDefinition) SetConfigurationHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleCONFIGURATION
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.ConfigurationHandler = handle
}

func (h *DefaultSmartAppDefinition) SetOAuthCallbackHandler(handler Handler) {
	handle := &DefaultPredicateHandler{
		Tester: func(params *SmartAppParams) bool {
			request := params.Request
			return request != nil && request.Lifecycle == smartapp.AppLifecycleOAUTHCALLBACK
		},
		Handler: func(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
			return handler(params)
		},
	}
	h.OAuthCallbackHandler = handle
}

func (h *DefaultSmartAppDefinition) AddHandler(handler PredicateHandler) {
	if h.AdditionalHandlers == nil {
		h.AdditionalHandlers = make([]PredicateHandler, 0, 3)
	}
	h.AdditionalHandlers = append(h.AdditionalHandlers, handler)
}

func (h *DefaultSmartAppDefinition) BuildChain() []PredicateHandler {
	chain := make([]PredicateHandler, 0, 8)

	if h.InstallHandler == nil {
		h.InstallHandler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleINSTALL,
		}
	}

	if h.UpdateHandler == nil {
		h.UpdateHandler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleUPDATE,
		}
	}

	if h.UninstallHandler == nil {
		h.UninstallHandler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleUNINSTALL,
		}
	}

	if h.EventHandler == nil {
		h.EventHandler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleEVENT,
		}
	}

	if h.PingHandler == nil {
		h.PingHandler = &PingHandler{}
	}

	if h.OAuthCallbackHandler == nil {
		h.OAuthCallbackHandler = &NoopLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleOAUTHCALLBACK,
		}
	}

	if h.ConfigurationHandler == nil {
		log.Panicf("Missing required lifecycle handler for lifecycle=%s", smartapp.AppLifecycleCONFIGURATION)
	}

	if len(h.AdditionalHandlers) > 0 {
		chain = append(chain, h.AdditionalHandlers...)
	}
	chain = append(chain, h.EventHandler)
	chain = append(chain, h.ConfigurationHandler)
	chain = append(chain, h.InstallHandler)
	chain = append(chain, h.UpdateHandler)
	chain = append(chain, h.UninstallHandler)
	chain = append(chain, h.PingHandler)
	chain = append(chain, h.OAuthCallbackHandler)
	chain = append(chain, &NotFoundHandler{})

	return chain
}
