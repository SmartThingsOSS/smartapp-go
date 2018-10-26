package handlers

import (
	"errors"
	"github.com/SmartThingsOSS/smartapp-go/examples/c2c-connector/internal/logging"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartapp"
	"github.com/SmartThingsOSS/smartapp-go/pkg/smartappcore"
	log "github.com/sirupsen/logrus"
)

var (
	UnsupportedConfigurationPhaseError = errors.New("unsupported_configuration_phase")
)

type ConfigHandler struct {
}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{}
}

func (h *ConfigHandler) Handler() smartappcore.Handler {
	return func(params *smartappcore.SmartAppParams) (*smartapp.ExecutionResponse, error) {
		logger := logging.Logger(params.Context)
		request := params.Request
		configData := request.ConfigurationData
		logger.WithFields(log.Fields{
			"installedAppId": configData.InstalledAppID,
			"phase":          configData.Phase,
		}).Debug("Config lifecycle...")

		switch configData.Phase {
		case smartapp.ConfigurationPhaseINITIALIZE:
			return &smartapp.ExecutionResponse{
				StatusCode: 200,
				ConfigurationData: &smartapp.ConfigurationResponseData{
					Initialize: &smartapp.InitializeSetting{
						Setting: smartapp.Setting{
							Name:        "Sample C2C SmartApp",
							Description: "Simple no-op C2C SmartApp",
							ID:          "app",
						},
						FirstPageID: "1",
						Permissions: []string{"i:deviceprofiles", "r:devices:*"},
					},
				},
			}, nil
		case smartapp.ConfigurationPhasePAGE:
			return &smartapp.ExecutionResponse{
				StatusCode: 200,
				ConfigurationData: &smartapp.ConfigurationResponseData{
					Page: &smartapp.Page{
						PageID:   "1",
						Name:     "Simple Automation app for certification testing",
						Complete: boolPointer(true),
						Sections: []*smartapp.Section{
							{
								Settings: []*smartapp.SectionSetting{
									{
										Setting: smartapp.Setting{
											Name:        "Certification Test C2C Connector",
											Description: "This SmartApp is a no-op automation that exists to verify the certification flow.",
											ID:          "myParagraphSetting",
										},
										Type: smartapp.SettingTypePARAGRAPH,
									},
								},
							},
						},
					},
				},
			}, nil
		default:
			return &smartapp.ExecutionResponse{}, UnsupportedConfigurationPhaseError
		}
	}
}

func boolPointer(val bool) *bool {
	return &val
}
