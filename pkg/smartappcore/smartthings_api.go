package smartappcore

import (
	"bytes"
	"context"
	apiclient "github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/devices"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/locations"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/scenes"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/schedules"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/client/subscriptions"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartthings/models"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/runtime/logger"
	"github.com/go-openapi/strfmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SmartThingsApi interface {
	CreateSchedule(params *schedules.CreateScheduleParams) (*models.Schedule, error)
	SaveSubscription(params *subscriptions.SaveSubscriptionParams) error
	InstallDevice(params *devices.InstallDeviceParams) (*models.Device, error)
	CreateDeviceEvents(params *devices.CreateDeviceEventsParams) error
	ExecuteDeviceCommands(params *devices.ExecuteDeviceCommandsParams) error
	GetLocation(params *locations.GetLocationParams) (*models.Location, error)
	ExecuteScene(params *scenes.ExecuteSceneParams) error
}

type SmartThingsApiParams struct {
	Host    string
	Schemes []string
	Debug   bool
	Logger  logger.Logger
}

type DefaultSmartThingsApi struct {
	Client         *apiclient.SmartThings
	RequestTimeout time.Duration
	Logger         logger.Logger
}

func NewSmartThingsApi(params *SmartThingsApiParams) SmartThingsApi {
	initParams(params)
	transport := httptransport.New(params.Host, "", params.Schemes)
	transport.SetLogger(params.Logger)
	transport.Transport = RequestInterceptor(transport.Transport)

	if params.Debug {
		transport.Transport = DebugInterceptor(params.Logger, transport.Transport)
	}

	client := apiclient.New(transport, strfmt.Default)
	return &DefaultSmartThingsApi{Client: client}
}

func (s *DefaultSmartThingsApi) CreateSchedule(params *schedules.CreateScheduleParams) (*models.Schedule, error) {
	ok, err := s.Client.Schedules.CreateSchedule(params, httptransport.BearerToken(params.Authorization))

	if err != nil {
		return nil, err
	}
	return ok.Payload, nil
}

func (s *DefaultSmartThingsApi) SaveSubscription(params *subscriptions.SaveSubscriptionParams) error {
	_, err := s.Client.Subscriptions.SaveSubscription(params, httptransport.BearerToken(params.Authorization))

	return err
}

func (s *DefaultSmartThingsApi) InstallDevice(params *devices.InstallDeviceParams) (*models.Device, error) {
	ok, err := s.Client.Devices.InstallDevice(params, httptransport.BearerToken(params.Authorization))

	if err != nil {
		return nil, err
	}

	return ok.Payload, nil
}

func (s *DefaultSmartThingsApi) CreateDeviceEvents(params *devices.CreateDeviceEventsParams) error {
	_, err := s.Client.Devices.CreateDeviceEvents(params, httptransport.BearerToken(params.Authorization))

	return err
}

func (s *DefaultSmartThingsApi) ExecuteDeviceCommands(params *devices.ExecuteDeviceCommandsParams) error {
	_, err := s.Client.Devices.ExecuteDeviceCommands(params, httptransport.BearerToken(params.Authorization))

	return err
}

func (s *DefaultSmartThingsApi) GetLocation(params *locations.GetLocationParams) (*models.Location, error) {
	ok, err := s.Client.Locations.GetLocation(params, httptransport.BearerToken(params.Authorization))

	if err != nil {
		return nil, err
	}

	return ok.Payload, nil
}

func (s *DefaultSmartThingsApi) ExecuteScene(params *scenes.ExecuteSceneParams) error {
	_, err := s.Client.Scenes.ExecuteScene(params, httptransport.BearerToken(params.Authorization))

	return err
}

func getRequestHeaders(ctx context.Context) map[string]string {
	headers := make(map[string]string)
	correlationId := ctx.Value(CorrelationKey)
	if correlationId != nil && correlationId != "" {
		headers[CorrelationHeader] = correlationId.(string)
	}
	headers["Accept"] = MediaTypeJsonV1

	return headers
}

func RequestInterceptor(original http.RoundTripper) http.RoundTripper {
	return roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		headers := getRequestHeaders(request.Context())
		for k, v := range headers {
			request.Header.Set(k, v)
		}
		return original.RoundTrip(request)
	})
}

func DebugInterceptor(log logger.Logger, original http.RoundTripper) http.RoundTripper {
	return roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		// Read Body
		requestBody, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Printf("Unable to read HTTP request body. error=%q", err)
		}

		// Reset body as it was cleared due to above read.
		request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		url := request.URL.String()
		method := request.Method

		log.Printf(
			"HTTP Request %s=%s, headers=%s, method=%s, url=%q, body=%q",
			CorrelationKey, request.Context().Value(CorrelationKey),
			request.Header, method, url, toString(requestBody),
		)

		start := time.Now()

		response, err := original.RoundTrip(request)
		if err != nil {
			return response, err
		}

		end := time.Now()
		latency := end.Sub(start)

		// Read Body
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		// Reset body as it was cleared due to above read.
		response.Body = ioutil.NopCloser(bytes.NewBuffer(responseBody))

		log.Printf(
			"HTTP Response %s=%s, status=%d, latency=%s, method=%s, url=%q, body=%q",
			CorrelationKey, request.Context().Value(CorrelationKey),
			response.StatusCode, latency, method, url, toString(responseBody),
		)

		return response, err
	})
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func toString(data []byte) string {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(bytes.NewBuffer(data)); err != nil {
		return ""
	}
	return buf.String()
}

func initParams(params *SmartThingsApiParams) {
	if params.Host == "" {
		params.Host = SmartThingsHost
	}

	if len(params.Schemes) == 0 {
		params.Schemes = []string{"https"}
	}

	if params.Logger == nil {
		params.Logger = logger.StandardLogger{}
	}
}
