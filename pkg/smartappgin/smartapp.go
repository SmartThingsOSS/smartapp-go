package smartappgin

import (
	"crypto/rsa"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RequestInterceptor func(params *smartappcore.SmartAppParams)
type ResponseInterceptor func(response *smartapp.ExecutionResponse, err error)

type SmartApp interface {
	Handler() gin.HandlerFunc
	SetPublicKey(publicKey *rsa.PublicKey)
	SetRequestInterceptor(interceptor RequestInterceptor)
	SetResponseInterceptor(interceptor ResponseInterceptor)
}

type DefaultSmartApp struct {
	App                 smartappcore.SmartApp
	Authenticator       *smartappcore.Authenticator
	RequestInterceptor  RequestInterceptor
	ResponseInterceptor ResponseInterceptor
}

func NewSmartApp(definition smartappcore.SmartAppDefinition) SmartApp {
	return &DefaultSmartApp{
		App:                 smartappcore.NewSmartApp(definition),
		RequestInterceptor:  func(params *smartappcore.SmartAppParams) {},
		ResponseInterceptor: func(response *smartapp.ExecutionResponse, err error) {},
	}
}

func (a *DefaultSmartApp) SetPublicKey(publicKey *rsa.PublicKey) {
	if publicKey != nil {
		a.Authenticator = smartappcore.NewAuthenticator(publicKey)
	}
}

func (a *DefaultSmartApp) SetRequestInterceptor(interceptor RequestInterceptor) {
	if interceptor != nil {
		a.RequestInterceptor = interceptor
	}
}

func (a *DefaultSmartApp) SetResponseInterceptor(interceptor ResponseInterceptor) {
	if interceptor != nil {
		a.ResponseInterceptor = interceptor
	}
}

func (a *DefaultSmartApp) Handler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request smartapp.ExecutionRequest
		var response *smartapp.ExecutionResponse
		var err error

		if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
			ctx.AbortWithError(422, err)
		}

		params := &smartappcore.SmartAppParams{
			Request: &request,
			Context: ctx,
		}

		a.RequestInterceptor(params)

		if a.Authenticator != nil && request.Lifecycle != smartapp.AppLifecyclePING {
			err := a.Authenticator.Verifier.Verify(ctx.Request)
			if err != nil {
				ctx.Status(401)
				return
			}
		}

		response, err = a.App.Handle(params)

		a.ResponseInterceptor(response, err)

		if err != nil {
			ctx.Error(err)
			return
		}

		statusCode := 500
		if response.StatusCode > 0 && response.StatusCode < 600 {
			statusCode = int(response.StatusCode)
		}

		ctx.JSON(statusCode, response)
	}
}
