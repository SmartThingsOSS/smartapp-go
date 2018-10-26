package smartappcore

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
)

type TestLifecycleHandler struct {
	Lifecycle smartapp.AppLifecycle
	Count     int
	Error     error
}

func (h *TestLifecycleHandler) Test(params *SmartAppParams) bool {
	return h.Lifecycle == params.Request.Lifecycle
}

func (h *TestLifecycleHandler) Handle(params *SmartAppParams) (*smartapp.ExecutionResponse, error) {
	h.Count = h.Count + 1
	return &smartapp.ExecutionResponse{StatusCode: 200}, h.Error
}

var _ = Describe("NotFoundHandler", func() {

	var (
		app                        SmartApp
		installHandler             *TestLifecycleHandler
		updateHandler              *TestLifecycleHandler
		configHandler              *TestLifecycleHandler
		uninstallHandler           *TestLifecycleHandler
		oauthHandler               *TestLifecycleHandler
		eventHandler               EventHandler
		deviceCommandsEventHandler *TestEventHandler
	)

	BeforeEach(func() {
		installHandler = &TestLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleINSTALL,
		}
		updateHandler = &TestLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleUPDATE,
		}
		configHandler = &TestLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleCONFIGURATION,
		}
		uninstallHandler = &TestLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleUNINSTALL,
		}
		oauthHandler = &TestLifecycleHandler{
			Lifecycle: smartapp.AppLifecycleUNINSTALL,
		}
		eventHandler = NewEventHandler()
		deviceCommandsEventHandler = &TestEventHandler{
			Value: true,
		}
		eventHandler.AddEventHandler(deviceCommandsEventHandler)
		definition := NewSmartAppDefinition()
		definition.SetInstallHandler(installHandler.Handle)
		definition.SetUpdateHandler(updateHandler.Handle)
		definition.SetConfigurationHandler(configHandler.Handle)
		definition.SetUninstallHandler(uninstallHandler.Handle)
		definition.SetOAuthCallbackHandler(oauthHandler.Handle)
		definition.SetEventHandler(eventHandler.BuildHandler())
		app = NewSmartApp(definition)
	})

	Context("With an instance of the default SmartApp implementation", func() {
		It("should process INSTALL lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle:   smartapp.AppLifecycleINSTALL,
					InstallData: &smartapp.InstallData{},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(1))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
		})

		It("should process UPDATE lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle:  smartapp.AppLifecycleUPDATE,
					UpdateData: &smartapp.UpdateData{},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(1))
			Expect(configHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
		})

		It("should process CONFIGURATION lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle:         smartapp.AppLifecycleCONFIGURATION,
					ConfigurationData: &smartapp.ConfigurationData{},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(1))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
		})

		It("should process UNINSTALL lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle:     smartapp.AppLifecycleUNINSTALL,
					UninstallData: &smartapp.UninstallData{},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(1))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
		})

		It("should process PING lifecycle", func() {
			aChallenge := "challenge"
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecyclePING,
					PingData: &smartapp.PingData{
						Challenge: &aChallenge,
					},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(response.PingData.Challenge).To(Equal(aChallenge))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
		})

		It("should process EVENT lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle: smartapp.AppLifecycleEVENT,
					EventData: &smartapp.EventData{
						Events: []*smartapp.Event{
							{
								EventType: smartapp.EventTypeDEVICECOMMANDSEVENT,
								DeviceCommandsEvent: &smartapp.DeviceCommandsEvent{
									DeviceID: "deviceId",
								},
							},
						},
					},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(1))
		})

		It("should process OAUTH_CALLBACK lifecycle", func() {
			response, err := app.Handle(&SmartAppParams{
				Request: &smartapp.ExecutionRequest{
					Lifecycle:         smartapp.AppLifecycleOAUTHCALLBACK,
					OauthCallbackData: &smartapp.OAuthCallbackData{},
				},
			})

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(int64(200)))
			Expect(installHandler.Count).To(Equal(0))
			Expect(updateHandler.Count).To(Equal(0))
			Expect(configHandler.Count).To(Equal(0))
			Expect(uninstallHandler.Count).To(Equal(0))
			Expect(deviceCommandsEventHandler.Count).To(Equal(0))
			Expect(oauthHandler.Count).To(Equal(1))
		})

	})
})
